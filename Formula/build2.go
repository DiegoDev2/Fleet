package main

// Build2 - C/C++ Build Toolchain
// Homepage: https://build2.org

import (
	"fmt"
	
	"os/exec"
)

func installBuild2() {
	// Método 1: Descargar y extraer .tar.gz
	build2_tar_url := "https://download.build2.org/0.17.0/build2-toolchain-0.17.0.tar.xz"
	build2_cmd_tar := exec.Command("curl", "-L", build2_tar_url, "-o", "package.tar.gz")
	err := build2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	build2_zip_url := "https://download.build2.org/0.17.0/build2-toolchain-0.17.0.tar.xz"
	build2_cmd_zip := exec.Command("curl", "-L", build2_zip_url, "-o", "package.zip")
	err = build2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	build2_bin_url := "https://download.build2.org/0.17.0/build2-toolchain-0.17.0.tar.xz"
	build2_cmd_bin := exec.Command("curl", "-L", build2_bin_url, "-o", "binary.bin")
	err = build2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	build2_src_url := "https://download.build2.org/0.17.0/build2-toolchain-0.17.0.tar.xz"
	build2_cmd_src := exec.Command("curl", "-L", build2_src_url, "-o", "source.tar.gz")
	err = build2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	build2_cmd_direct := exec.Command("./binary")
	err = build2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
