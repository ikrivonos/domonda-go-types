package email

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttachment_MarshalJSON(t *testing.T) {
	j, err := json.Marshal(Attachment{
		PartID:      "PartID",
		ContentID:   "ContentID",
		ContentType: "ContentType",
		Filename:    "FileName",
		Content:     []byte("FileData"),
	})
	require.NoError(t, err, "json.Marshal")
	require.Equal(t, `{"PartID":"PartID","ContentID":"ContentID","ContentType":"ContentType","Filename":"FileName","Content":"RmlsZURhdGE="}`, string(j))
}
