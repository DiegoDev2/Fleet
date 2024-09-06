package main

// BuildpulseTestReporter - Connect your CI to BuildPulse to detect, track, and rank flaky tests
// Homepage: https://buildpulse.io

import (
	"fmt"
	
	"os/exec"
)

func installBuildpulseTestReporter() {
	// Método 1: Descargar y extraer .tar.gz
	buildpulsetestreporter_tar_url := "https://github.com/buildpulse/test-reporter/archive/refs/tags/v0.28.0.tar.gz"
	buildpulsetestreporter_cmd_tar := exec.Command("curl", "-L", buildpulsetestreporter_tar_url, "-o", "package.tar.gz")
	err := buildpulsetestreporter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	buildpulsetestreporter_zip_url := "https://github.com/buildpulse/test-reporter/archive/refs/tags/v0.28.0.zip"
	buildpulsetestreporter_cmd_zip := exec.Command("curl", "-L", buildpulsetestreporter_zip_url, "-o", "package.zip")
	err = buildpulsetestreporter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	buildpulsetestreporter_bin_url := "https://github.com/buildpulse/test-reporter/archive/refs/tags/v0.28.0.bin"
	buildpulsetestreporter_cmd_bin := exec.Command("curl", "-L", buildpulsetestreporter_bin_url, "-o", "binary.bin")
	err = buildpulsetestreporter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	buildpulsetestreporter_src_url := "https://github.com/buildpulse/test-reporter/archive/refs/tags/v0.28.0.src.tar.gz"
	buildpulsetestreporter_cmd_src := exec.Command("curl", "-L", buildpulsetestreporter_src_url, "-o", "source.tar.gz")
	err = buildpulsetestreporter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	buildpulsetestreporter_cmd_direct := exec.Command("./binary")
	err = buildpulsetestreporter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
