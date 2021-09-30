package parse

import (
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func JSONtoProto(rawJSON []byte, data protoreflect.ProtoMessage) ([]byte, error) {
	err := protojson.Unmarshal(rawJSON, data)
	if err != nil {
		return nil, fmt.Errorf("error unmashalling raw json into protobuf: %w", err)
	}

	output, err := proto.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling protobuf to wire format: %w", err)
	}

	return output, nil
}
func PipeJSONtoProto(data protoreflect.ProtoMessage) error {
	rawJSON, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("error reading data from stdin: %w", err)
	}

	output, err := JSONtoProto(rawJSON, data)
	if err != nil {
		return fmt.Errorf("error converting json to protobuf: %w", err)
	}

	_, err = os.Stdout.Write(output)
	if err != nil {
		return fmt.Errorf("error writing protobuf wire format to stdout: %w", err)
	}
	return nil
}
