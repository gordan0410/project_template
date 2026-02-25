## 2026-02-25

### 重構目錄結構以符合 CLAUDE.md 規範

**修改檔案：**
- `cmd/wire.go`：將已註解的 `//api.ProvideSet` 更新為 `//client.ProvideSet`

**新增目錄/檔案：**
- `internal/constant/` + `.gitkeep`：新增常數目錄，取代反模式的 `enum` 命名
- `internal/client/provider.go`：新增外部 API 存取層佔位符，package 名稱為 `client`

**移除：**
- `internal/enum/`：空目錄，已重命名為 `internal/constant/`
- `internal/api/`：空的 Wire provider 目錄，已重命名為 `internal/client/`
