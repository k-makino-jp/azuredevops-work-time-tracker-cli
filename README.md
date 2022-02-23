# AzureDevOps Work Time Tracker CLI

本リポジトリは、[Azure DevOps](https://azure.microsoft.com/ja-jp/services/devops/) の Boards で管理されるチケットの作業時間追跡 CLI ツールを提供します。

## 前提条件

本 CLI ツールを利用するためには、以下のソフトウェアがインストールされている必要があります。

* Golang 1.16 以上

また、以下のスコープの [Personal Access Token](https://docs.microsoft.com/ja-jp/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops&tabs=preview-page#create-a-pat) を作成する必要があります。

* Work Items: Read

## 提供機能

本 CLI ツールは、以下の機能を提供します。

* [Updates - List](https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/updates/list?view=azure-devops-rest-6.0) API で WorkItem の New、Active または Closed の 時間を集計します。

## 実行方法

本 CLI ツールは、以下の手順で実行します。

1. 本プロジェクトのディレクトリに移動します。
   ```sh
   $ cd azuredevops-work-time-tracker-cli
   ```
1. カレントディレクトリーの `config.json` を編集します。
   * フォーマット
      ```json
      {
         "organization":"your-organization-name" ,
         "project":"your-project-name"
      }
      ```
   * 例
       ```json
       {
           "organization":"worktimetracker-org" ,
           "project":"worktimetracker-prj"
       }
       ```
2. 以下のコマンドを実行します。
   ```sh
   $ go run main.go get -i <WorkItem ID> -p "<Personal Access Token>" -s "<StatusA,StatusB>"

   [Example]
   $ go run main.go get -i 100 -p "personalaccesstoken" -s "New,Active"

   # Target State: New

   ## Histories

   | Start Time           | End Time             | Spend Time [hours]   |
   | -------------------- | ---------------------| -------------------- |
   | 2022-02-21T00:35:43Z | 2022-02-21T00:35:47Z | 00.00                |
   | 2022-02-21T00:35:49Z | 2022-02-21T00:36:59Z | 00.02                |
   | 2022-02-21T00:37:03Z | 2022-02-21T00:37:05Z | 00.00                |
   | 2022-02-21T00:37:09Z | 2022-02-21T02:05:54Z | 01.48                |

   ## Total Spend Time

   * 01.50 [hours]

   # Target State: Active

   ## Histories

   | Start Time           | End Time             | Spend Time [hours]   |
   | -------------------- | ---------------------| -------------------- |
   | 2022-02-21T00:35:47Z | 2022-02-21T00:35:49Z | 00.00                |
   | 2022-02-21T00:36:59Z | 2022-02-21T00:37:03Z | 00.00                |
   | 2022-02-21T00:37:05Z | 2022-02-21T00:37:06Z | 00.00                |
   | 2022-02-21T00:37:07Z | 2022-02-21T00:37:09Z | 00.00                |
   | 2022-02-21T02:05:54Z | 2022-02-21T05:00:06Z | 02.90                |

   ## Total Spend Time

   * 02.91 [hours]
   ```
