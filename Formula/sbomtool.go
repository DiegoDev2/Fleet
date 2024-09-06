package main

// SbomTool - Scalable and enterprise ready tool to create SBOMs for any variety of artifacts
// Homepage: https://github.com/microsoft/sbom-tool

import (
	"fmt"
	
	"os/exec"
)

func installSbomTool() {
	// Método 1: Descargar y extraer .tar.gz
	sbomtool_tar_url := "https://github.com/microsoft/sbom-tool/archive/refs/tags/v2.2.7.tar.gz"
	sbomtool_cmd_tar := exec.Command("curl", "-L", sbomtool_tar_url, "-o", "package.tar.gz")
	err := sbomtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sbomtool_zip_url := "https://github.com/microsoft/sbom-tool/archive/refs/tags/v2.2.7.zip"
	sbomtool_cmd_zip := exec.Command("curl", "-L", sbomtool_zip_url, "-o", "package.zip")
	err = sbomtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sbomtool_bin_url := "https://github.com/microsoft/sbom-tool/archive/refs/tags/v2.2.7.bin"
	sbomtool_cmd_bin := exec.Command("curl", "-L", sbomtool_bin_url, "-o", "binary.bin")
	err = sbomtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sbomtool_src_url := "https://github.com/microsoft/sbom-tool/archive/refs/tags/v2.2.7.src.tar.gz"
	sbomtool_cmd_src := exec.Command("curl", "-L", sbomtool_src_url, "-o", "source.tar.gz")
	err = sbomtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sbomtool_cmd_direct := exec.Command("./binary")
	err = sbomtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet")
	exec.Command("latte", "install", "dotnet").Run()
}
