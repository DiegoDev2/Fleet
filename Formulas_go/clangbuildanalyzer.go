package main

// ClangBuildAnalyzer - Tool to analyze compilation time
// Homepage: https://github.com/aras-p/ClangBuildAnalyzer

import (
	"fmt"
	
	"os/exec"
)

func installClangBuildAnalyzer() {
	// Método 1: Descargar y extraer .tar.gz
	clangbuildanalyzer_tar_url := "https://github.com/aras-p/ClangBuildAnalyzer/archive/refs/tags/v1.5.0.tar.gz"
	clangbuildanalyzer_cmd_tar := exec.Command("curl", "-L", clangbuildanalyzer_tar_url, "-o", "package.tar.gz")
	err := clangbuildanalyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clangbuildanalyzer_zip_url := "https://github.com/aras-p/ClangBuildAnalyzer/archive/refs/tags/v1.5.0.zip"
	clangbuildanalyzer_cmd_zip := exec.Command("curl", "-L", clangbuildanalyzer_zip_url, "-o", "package.zip")
	err = clangbuildanalyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clangbuildanalyzer_bin_url := "https://github.com/aras-p/ClangBuildAnalyzer/archive/refs/tags/v1.5.0.bin"
	clangbuildanalyzer_cmd_bin := exec.Command("curl", "-L", clangbuildanalyzer_bin_url, "-o", "binary.bin")
	err = clangbuildanalyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clangbuildanalyzer_src_url := "https://github.com/aras-p/ClangBuildAnalyzer/archive/refs/tags/v1.5.0.src.tar.gz"
	clangbuildanalyzer_cmd_src := exec.Command("curl", "-L", clangbuildanalyzer_src_url, "-o", "source.tar.gz")
	err = clangbuildanalyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clangbuildanalyzer_cmd_direct := exec.Command("./binary")
	err = clangbuildanalyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
