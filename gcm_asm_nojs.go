// +build !js

package qtls

import (
	"golang.org/x/sys/cpu"
	"os"
	"strings"
)

func getHasGCMAsm() bool {
	// Check the cpu flags for each platform that has optimized GCM implementations.
	// Worst case, these variables will just all be false.
	var (
		hasGCMAsmAMD64 = cpu.X86.HasAES && cpu.X86.HasPCLMULQDQ
		hasGCMAsmARM64 = cpu.ARM64.HasAES && cpu.ARM64.HasPMULL
		// Keep in sync with crypto/aes/cipher_s390x.go.
		// TODO: check for s390
		// hasGCMAsmS390X = cpu.S390X.HasAES && cpu.S390X.HasAESCBC && cpu.S390X.HasAESCTR && (cpu.S390X.HasGHASH || cpu.S390X.HasAESGCM)
		hasGCMAsmS390X = false

		hasGCMAsm = hasGCMAsmAMD64 || hasGCMAsmARM64 || hasGCMAsmS390X
	)

	// x/sys/cpu does not respect GODEBUG=cpu.all=off. As a workaround,
	// check it here. See https://github.com/golang/go/issues/33963
	if strings.Contains(os.Getenv("GODEBUG"), "cpu.all=off") {
		hasGCMAsm = false
	}

	return hasGCMAsm
}
