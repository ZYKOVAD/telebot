package updates

import (
	"encoding/json"
	"fmt"
	"net/http"
	"telegram_bot/internal/common"
)

func GetUpdates(offset int) ([]common.Update, error) {
	url := fmt.Sprintf(common.TelegramAPI, common.Token, "getUpdates")
	if offset > 0 {
		url += fmt.Sprintf("?offset=%d", offset)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var updateResp common.UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		return nil, err
	}
	return updateResp.Result, nil
}
