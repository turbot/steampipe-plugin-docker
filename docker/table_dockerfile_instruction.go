package docker

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/pkg/errors"

	filehelpers "github.com/turbot/go-kit/files"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDockerfileInstruction(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "dockerfile_instruction",
		Description: "List all instructions from the Dockerfile.",
		List: &plugin.ListConfig{
			ParentHydrate: dockerfileList,
			Hydrate:       listDockerfileInstruction,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "path", Type: proto.ColumnType_STRING, Description: "Full path of the file."},
			{Name: "stage", Type: proto.ColumnType_STRING, Description: "Stage name in the Dockerfile, defaults to the stage number."},
			{Name: "instruction", Type: proto.ColumnType_STRING, Description: "Command name in lowercase form, e.g. from, env, run, etc."},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "Command data, parsed into a convenient format for each command type."},
			// Other columns
			{Name: "args", Type: proto.ColumnType_JSON, Description: "Array of arguments passed to the command."},
			{Name: "end_line", Type: proto.ColumnType_INT, Description: "Last line number of this cmd in the file."},
			{Name: "flags", Type: proto.ColumnType_JSON, Description: "Flags passed to the command."},
			{Name: "prev_comment", Type: proto.ColumnType_JSON, Transform: transform.FromField("PrevComment"), Description: "Comment above the command in the Dockerfile."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "Full original source code of the cmd."},
			{Name: "stage_number", Type: proto.ColumnType_INT, Description: "Stage number in the Dockerfile, starting at zero."},
			{Name: "start_line", Type: proto.ColumnType_INT, Description: "First line number of this cmd in the file."},
			{Name: "sub_instruction", Type: proto.ColumnType_STRING, Description: "Sub command name in lowercase form, e.g. set to 'run' for 'onbuild run ...'."},
		},
	}
}

// Command is the struct for each dockerfile command
type Command struct {
	Path           string
	Stage          string
	StageNumber    int
	Instruction    string
	SubInstruction string
	Flags          []string
	Args           []string
	PrevComment    []string
	Source         string
	StartLine      int
	EndLine        int
	Data           interface{}
}

type filePath struct {
	Path string
}

type nameValuePair struct {
	Name  string  `json:"name"`
	Value *string `json:"value,omitempty"`
}

type argInstructionData struct {
	Args []nameValuePair `json:"args"`
}

type copyInstructionData struct {
	Sources []string `json:"sources"`
	Dest    string   `json:"dest"`
	Chown   string   `json:"chown,omitempty"`
	Chmod   string   `json:"chmod,omitempty"`
}

type exposeInstructionData struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type fromInstructionData struct {
	Image     string `json:"image"`
	Tag       string `json:"tag,omitempty"`
	Digest    string `json:"digest,omitempty"`
	StageName string `json:"stage_name,omitempty"`
}

type runInstructionData struct {
	PrependShell bool     `json:"prepend_shell,omitempty"`
	Commands     []string `json:"commands"`
}

type userInstructionData struct {
	User  string `json:"user"`
	Group string `json:"group,omitempty"`
}

type volumeInstructionData struct {
	Volumes []string `json:"volumes"`
}

type workdirInstructionData struct {
	Path string `json:"path"`
}

func dockerfileList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// #1 - Path via qual

	// If the path was requested through qualifier then match it exactly. Globs
	// are not supported in this context since the output value for the column
	// will never match the requested value.
	quals := d.EqualsQuals
	if quals["path"] != nil {
		d.StreamListItem(ctx, filePath{Path: quals["path"].GetStringValue()})
		return nil, nil
	}

	// #2 - Glob paths in config

	// Fail if no paths are specified
	dockerConfig := GetConfig(d.Connection)

	plugin.Logger(ctx).Warn("paths parameter is deprecated and will be removed after 31st August 2023, please use dockerfile_paths instead.")
	if dockerConfig.Paths == nil && dockerConfig.DockerfilePaths == nil {
		return nil, errors.New("dockerfile_paths must be configured")
	}

	// Gather file path matches for the glob
	var matches []string
	var dPath []string

	if dockerConfig.Paths != nil {
		dPath = append(dPath, dockerConfig.Paths...)
	} else {
		dPath = append(dPath, dockerConfig.DockerfilePaths...)
	}
	for _, i := range dPath {
		// List the files in the given source directory
		files, err := d.GetSourceFiles(i)
		if err != nil {
			// If the specified path is unavailable, then an empty row should populate
			plugin.Logger(ctx).Error("dockerfile_instruction.dockerfileList", "get_source_files_error", err)
			if strings.Contains(err.Error(), "failed to get directory specified by the source") {
				return nil, nil
			}
			return nil, err
		}
		plugin.Logger(ctx).Warn("dockerfileList", "source", i, "files", files)
		matches = append(matches, files...)
	}

	// Sanitize the matches to ignore the directories
	for _, i := range matches {

		// Ignore directories
		if filehelpers.DirectoryExists(i) {
			continue
		}
		d.StreamListItem(ctx, filePath{Path: i})
	}

	return nil, nil
}

func listDockerfileInstruction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath)

	reader, err := os.Open(path.Path)
	if err != nil {
		// Could not open the file, so log and ignore
		plugin.Logger(ctx).Error("listDockerfileInstruction", "file_error", err, "path", path.Path)
		return nil, nil
	}

	parsed, err := parser.Parse(reader)
	if err != nil {
		// Could not parse the file, so log and ignore
		plugin.Logger(ctx).Error("listDockerfileInstruction", "parse_error", err, "path", path.Path)
		return nil, nil
	}

	stage := ""
	stageNumber := -1

	for _, i := range parsed.AST.Children {
		cmd := Command{
			Path:        path.Path,
			Instruction: i.Value,
			Source:      i.Original,
			Flags:       i.Flags,
			StartLine:   i.StartLine,
			EndLine:     i.EndLine,
			PrevComment: i.PrevComment,
		}

		if i.Next != nil && len(i.Next.Children) > 0 {
			child := i.Next.Children[0]
			cmd.SubInstruction = child.Value
			cmd.Args = append(cmd.Args, child.Value)
			for n := child.Next; n != nil; n = n.Next {
				cmd.Args = append(cmd.Args, n.Value)
			}
		}

		for n := i.Next; n != nil; n = n.Next {
			cmd.Args = append(cmd.Args, n.Value)
		}

		if i.Value == "from" {
			stageNumber++
			stage = fmt.Sprintf("%d", stageNumber)
			if cmd.Instruction == "from" && len(cmd.Args) >= 3 {
				if strings.ToLower(cmd.Args[1]) == "as" {
					stage = cmd.Args[2]
				}
			}
		}
		cmd.Stage = stage
		cmd.StageNumber = stageNumber

		instruction, err := instructions.ParseInstruction(i)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to parse file %s", path))
		}

		switch ic := instruction.(type) {
		case *instructions.AddCommand:
			data := copyInstructionData{
				Chmod:   ic.Chmod,
				Chown:   ic.Chown,
				Dest:    ic.SourcesAndDest.DestPath,
				Sources: ic.SourcesAndDest.SourcePaths,
			}
			cmd.Data = data
		case *instructions.ArgCommand:
			data := argInstructionData{}
			for _, i := range ic.Args {
				arg := nameValuePair{
					Name:  i.Key,
					Value: i.Value,
				}
				data.Args = append(data.Args, arg)
			}
			cmd.Data = data
		case *instructions.CmdCommand:
			data := runInstructionData{
				Commands: ic.CmdLine,
			}
			cmd.Data = data
		case *instructions.CopyCommand:
			data := copyInstructionData{
				Chmod:   ic.Chmod,
				Chown:   ic.Chown,
				Dest:    ic.SourcesAndDest.DestPath,
				Sources: ic.SourcesAndDest.SourcePaths,
			}
			cmd.Data = data
		case *instructions.EntrypointCommand:
			data := runInstructionData{
				Commands: ic.CmdLine,
			}
			cmd.Data = data
		case *instructions.EnvCommand:
			data := map[string]string{}
			for _, kv := range ic.Env {
				data[kv.Key] = kv.Value
			}
			cmd.Data = data
		case *instructions.ExposeCommand:
			data := []exposeInstructionData{}
			for _, p := range ic.Ports {
				parts := strings.Split(p, "/")
				iPort, err := strconv.Atoi(parts[0])
				if err != nil {
					plugin.Logger(ctx).Error("listDockerfileInstruction", "expose_data_parsing_error", err, "cmd", cmd)
					continue
				}
				ep := exposeInstructionData{
					Port:     iPort,
					Protocol: "tcp",
				}
				if len(parts) > 1 {
					ep.Protocol = parts[1]
				}
				data = append(data, ep)
			}
			cmd.Data = data
		case *instructions.HealthCheckCommand:
			cmd.Data = ic.Health
		case *instructions.LabelCommand:
			data := map[string]string{}
			for _, kv := range ic.Labels {
				data[kv.Key] = kv.Value
			}
			cmd.Data = data
		case *instructions.RunCommand:
			// NOTE: This is an approximate split of the commands only based on &&.
			// It does not do full parsing of the command so may be inaccurate if the
			// command includes && for other reasons (rare).
			re := regexp.MustCompile(`\s*&&\s*`)
			data := runInstructionData{
				PrependShell: ic.PrependShell,
				Commands:     re.Split(ic.CmdLine[0], -1),
			}
			cmd.Data = data
		case *instructions.UserCommand:
			data := userInstructionData{}
			parts := strings.Split(ic.User, ":")
			data.User = parts[0]
			if len(parts) >= 2 {
				data.Group = parts[1]
			}
			cmd.Data = data
		case *instructions.VolumeCommand:
			cmd.Data = volumeInstructionData{Volumes: ic.Volumes}
		case *instructions.WorkdirCommand:
			cmd.Data = workdirInstructionData{Path: ic.Path}
		}

		switch cmd.Instruction {
		case "from":
			data := fromInstructionData{}
			// Get the image and qualifier (if any)
			parts := strings.Split(cmd.Args[0], ":")
			if len(parts) >= 2 {
				data.Image = parts[0]
				data.Tag = parts[1]
			} else {
				parts := strings.Split(cmd.Args[0], "@")
				data.Image = parts[0]
				if len(parts) >= 2 {
					data.Digest = parts[1]
				}
			}
			// Get the stage name (if any)
			if len(cmd.Args) >= 3 {
				if strings.ToLower(cmd.Args[1]) == "as" {
					data.StageName = cmd.Args[2]
				}
			}
			cmd.Data = data
		}

		d.StreamListItem(ctx, cmd)
	}

	return nil, nil
}
