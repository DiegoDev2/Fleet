package main

// Vttest - Test compatibility of VT100-compatible terminals
// Homepage: https://invisible-island.net/vttest/

import (
	"fmt"
	
	"os/exec"
)

func installVttest() {
	// Método 1: Descargar y extraer .tar.gz
	vttest_tar_url := "https://invisible-mirror.net/archives/vttest/vttest-20240708.tgz"
	vttest_cmd_tar := exec.Command("curl", "-L", vttest_tar_url, "-o", "package.tar.gz")
	err := vttest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vttest_zip_url := "https://invisible-mirror.net/archives/vttest/vttest-20240708.tgz"
	vttest_cmd_zip := exec.Command("curl", "-L", vttest_zip_url, "-o", "package.zip")
	err = vttest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vttest_bin_url := "https://invisible-mirror.net/archives/vttest/vttest-20240708.tgz"
	vttest_cmd_bin := exec.Command("curl", "-L", vttest_bin_url, "-o", "binary.bin")
	err = vttest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vttest_src_url := "https://invisible-mirror.net/archives/vttest/vttest-20240708.tgz"
	vttest_cmd_src := exec.Command("curl", "-L", vttest_src_url, "-o", "source.tar.gz")
	err = vttest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vttest_cmd_direct := exec.Command("./binary")
	err = vttest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
