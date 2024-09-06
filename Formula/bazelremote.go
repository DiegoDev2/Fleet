package main

// BazelRemote - Remote cache for Bazel
// Homepage: https://github.com/buchgr/bazel-remote/

import (
	"fmt"
	
	"os/exec"
)

func installBazelRemote() {
	// Método 1: Descargar y extraer .tar.gz
	bazelremote_tar_url := "https://github.com/buchgr/bazel-remote/archive/refs/tags/v2.4.4.tar.gz"
	bazelremote_cmd_tar := exec.Command("curl", "-L", bazelremote_tar_url, "-o", "package.tar.gz")
	err := bazelremote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bazelremote_zip_url := "https://github.com/buchgr/bazel-remote/archive/refs/tags/v2.4.4.zip"
	bazelremote_cmd_zip := exec.Command("curl", "-L", bazelremote_zip_url, "-o", "package.zip")
	err = bazelremote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bazelremote_bin_url := "https://github.com/buchgr/bazel-remote/archive/refs/tags/v2.4.4.bin"
	bazelremote_cmd_bin := exec.Command("curl", "-L", bazelremote_bin_url, "-o", "binary.bin")
	err = bazelremote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bazelremote_src_url := "https://github.com/buchgr/bazel-remote/archive/refs/tags/v2.4.4.src.tar.gz"
	bazelremote_cmd_src := exec.Command("curl", "-L", bazelremote_src_url, "-o", "source.tar.gz")
	err = bazelremote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bazelremote_cmd_direct := exec.Command("./binary")
	err = bazelremote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
