package utils

type recoverFunc func()
type logErrFunc func(interface{})

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : panic recover
 * @param           {recoverFunc} goFunc
 * @param           {logErrFunc} logFunc
 * @return          {*}
 */
func RecoverFunc(goFunc recoverFunc, logFunc logErrFunc) {
	defer func(f recoverFunc) {
		if err := recover(); err != nil {
			go f()
			go logFunc(err)
		}
	}(goFunc)
	goFunc()
}
