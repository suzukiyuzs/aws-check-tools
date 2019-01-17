# aws-check-tools

## 概要
AWS環境の設定内容を確認する簡単なツールです。
コマンド実行後、OK/NGの数をカンマ区切りで出力します。

- コマンド実行例
```
$ コマンド
OKカウント数,NGカウント数
$
```

この実行結果をMackerelのサービスメトリックに投稿し、可視化/監視するような利用方法を想定しています。（=Mackerelに投げつけるため、ビルドしたコマンドをbash等でラップする想定）

### Check-CLB-RegisteredInstances
ELB(CLB)にEC2が紐付けられていればOKにカウント、1台もEC2が紐付けられていない場合はNGにカウント。

### Check-Cloudwatch-EC2Alarm
引数に指定したEC2メトリックのCloudwatchアラームのステータスをOKの場合はOKにカウント、ALARMの場合はNGにカウント。

### Check-EBS-AvailableVolume
availableのEBSボリュームをNGにカウント。available以外はOKにカウント。

### Check-EBS-SnapshotAmi
AMI作成時に取得したSnapshotで、基となるAMIが存在しない場合はNGにカウント。
AMIが存在する場合はOKにカウント。

### Check-EBS-SnapshotTag
引数に指定したタグがSnapshotに設定されていたらOKにカウント。設定されてなければNGにカウント。

### Check-EBS-VolumeDeleteFlag
EBSボリュームの「Delete on Terminate」が設定されてなかったらNGにカウント。
設定されていればOKにカウント。

### Check-EBS-VolumeTag
引数に指定したタグがEBSボリュームに設定されていたらOKにカウント。設定されてなければNGにカウント。

### Check-EC2-LaunchTime
EC2のLaunchTimeが引数に指定した時刻より古ければNGにカウント。若ければOKにカウント。

### Check-EC2-Tag
引数に指定したタグがEC2に設定されていたらOKにカウント。設定されてなければNGにカウント。

### Check-EIP-Association
関連付けされていないEIPをNGにカウント。関連済みはOKにカウント。


## ビルド
```
$ go get github.com/suzukiyuzs/aws-check-tools
$ cd $GOPATH/src/github.com/suzukiyuzs/aws-check-tools
$ cd << サブディレクトリ >>
$ go build .

```
※Check-EBS-SnapshotAmi/main.go, Check-EBS-SnapshotTag/main.goはビルド前にAWSアカウントIDの修正が必要。
※リージョンをハードコーディングしているので「ap-northeast-1」以外の場合はregionの値も変更が必要です。


## 実行例
以下は、AWSのAccess Key、Secret Access Keyは「aws configure」で設定済みの環境を想定した実行例です。

### Check-CLB-RegisteredInstances
```
$ ./Check-CLB-RegisteredInstances
20,4
$
```

### Check-Cloudwatch-EC2Alarm
```
$ ./Check-Cloudwatch-EC2Alarm CPUUtilization
10,1
$
```

### Check-EBS-AvailableVolume
```
$ ./Check-EBS-AvailableVolume
91,15
$
```

### Check-EBS-SnapshotAmi
```
$ ./Check-EBS-SnapshotAmi
120,88
$
```

### Check-EBS-SnapshotTag
```
./Check-EBS-SnapshotTag Name
6,210
$
```

### Check-EBS-VolumeDeleteFlag
```
$ Check-EBS-VolumeDeleteFlag
90,1
$
```

### Check-EBS-VolumeTag
```
./Check-EBS-VolumeTag Name
39,67
$
```

### Check-EC2-LaunchTime
```
$ ./Check-EC2-LaunchTime 2018-12-01T00:00:00
63,18
$
```

### Check-EC2-Tag
```
$ ./Check-EC2-Tag Stack
54,27
$
```

### Check-EIP-Association
```
$ ./Check-EIP-Association
55,1
$
```
