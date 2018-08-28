package config

import (
	"strings"

	"github.com/cihub/seelog"
)

// a default log cfg for seelog
const defaultLoggerCfg = `
<seelog>
    <outputs>
        <filter levels="trace">
            <console formatid="common"/>
        </filter>
        <filter levels="debug">
            <console formatid="coloredmagenta"/>
        </filter>
        <filter levels="info">
            <console formatid="coloredyellow"/>
        </filter>
        <filter levels="warn">
            <console formatid="coloredblue"/>
        </filter>
        <filter levels="error,critical">
            <splitter formatid="coloredred">
                <console/>
                <file path="./log/{ProgramName}_err.log"/>
            </splitter>
        </filter>
        <file formatid="common" path="./log/{ProgramName}.log"/>
    </outputs>
    <formats>
        <format id="coloredblue"  format="[%Date %Time] %EscM(34)[%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredred"  format="[%Date %Time] %EscM(31)[%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredgreen"  format="[%Date %Time] %EscM(32)[%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredyellow"  format="[%Date %Time] %EscM(33)[%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredcyan"  format="[%Date %Time] %EscM(36)[%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="coloredmagenta"  format="[%Date %Time] %EscM(35)[%LEV] [%File(%Line)] [%Func] %Msg%EscM(39)%n%EscM(0)"/>
        <format id="common"  format="[%Date %Time] [%LEV] [%File(%Line)] [%Func] %Msg%n"/>
        <format id="sentry"  format="%Msg%n"/>
    </formats>
</seelog>
`

// InitLogger init seelog
func InitLogger(name string, cfgPath string) {
	cfg := ""
	if cfgPath == "" {
		// use default logger
		if name == "" {
			InitLogger("common", "")
			seelog.Warn("log file no name, now is common")
			return
		}

		cfg = strings.Replace(defaultLoggerCfg, "{ProgramName}", name, -1)

		logger, err := seelog.LoggerFromConfigAsBytes([]byte(cfg))
		if err != nil {
			panic(err)
			return
		}
		seelog.ReplaceLogger(logger)
	} else {
		logger, err := seelog.LoggerFromConfigAsFile(cfgPath)
		if err != nil {
			InitLogger(name, "")
			seelog.Error("load log cfg err by %s", err.Error())
			return
		}
		seelog.ReplaceLogger(logger)
	}
}
