package main

// OsxCpuTemp - Outputs current CPU temperature for OSX
// Homepage: https://github.com/lavoiesl/osx-cpu-temp

import (
	"fmt"
	
	"os/exec"
)

func installOsxCpuTemp() {
	// Método 1: Descargar y extraer .tar.gz
	osxcputemp_tar_url := "https://github.com/lavoiesl/osx-cpu-temp/archive/refs/tags/1.1.0.tar.gz"
	osxcputemp_cmd_tar := exec.Command("curl", "-L", osxcputemp_tar_url, "-o", "package.tar.gz")
	err := osxcputemp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osxcputemp_zip_url := "https://github.com/lavoiesl/osx-cpu-temp/archive/refs/tags/1.1.0.zip"
	osxcputemp_cmd_zip := exec.Command("curl", "-L", osxcputemp_zip_url, "-o", "package.zip")
	err = osxcputemp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osxcputemp_bin_url := "https://github.com/lavoiesl/osx-cpu-temp/archive/refs/tags/1.1.0.bin"
	osxcputemp_cmd_bin := exec.Command("curl", "-L", osxcputemp_bin_url, "-o", "binary.bin")
	err = osxcputemp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osxcputemp_src_url := "https://github.com/lavoiesl/osx-cpu-temp/archive/refs/tags/1.1.0.src.tar.gz"
	osxcputemp_cmd_src := exec.Command("curl", "-L", osxcputemp_src_url, "-o", "source.tar.gz")
	err = osxcputemp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osxcputemp_cmd_direct := exec.Command("./binary")
	err = osxcputemp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
