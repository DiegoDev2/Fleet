package main

// LdFindCodeRefs - Build tool for sending feature flag code references to LaunchDarkly
// Homepage: https://github.com/launchdarkly/ld-find-code-refs

import (
	"fmt"
	
	"os/exec"
)

func installLdFindCodeRefs() {
	// Método 1: Descargar y extraer .tar.gz
	ldfindcoderefs_tar_url := "https://github.com/launchdarkly/ld-find-code-refs/archive/refs/tags/v2.12.0.tar.gz"
	ldfindcoderefs_cmd_tar := exec.Command("curl", "-L", ldfindcoderefs_tar_url, "-o", "package.tar.gz")
	err := ldfindcoderefs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ldfindcoderefs_zip_url := "https://github.com/launchdarkly/ld-find-code-refs/archive/refs/tags/v2.12.0.zip"
	ldfindcoderefs_cmd_zip := exec.Command("curl", "-L", ldfindcoderefs_zip_url, "-o", "package.zip")
	err = ldfindcoderefs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ldfindcoderefs_bin_url := "https://github.com/launchdarkly/ld-find-code-refs/archive/refs/tags/v2.12.0.bin"
	ldfindcoderefs_cmd_bin := exec.Command("curl", "-L", ldfindcoderefs_bin_url, "-o", "binary.bin")
	err = ldfindcoderefs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ldfindcoderefs_src_url := "https://github.com/launchdarkly/ld-find-code-refs/archive/refs/tags/v2.12.0.src.tar.gz"
	ldfindcoderefs_cmd_src := exec.Command("curl", "-L", ldfindcoderefs_src_url, "-o", "source.tar.gz")
	err = ldfindcoderefs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ldfindcoderefs_cmd_direct := exec.Command("./binary")
	err = ldfindcoderefs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
