package main

// SpdxSbomGenerator - Support CI generation of SBOMs via golang tooling
// Homepage: https://github.com/opensbom-generator/spdx-sbom-generator

import (
	"fmt"
	
	"os/exec"
)

func installSpdxSbomGenerator() {
	// Método 1: Descargar y extraer .tar.gz
	spdxsbomgenerator_tar_url := "https://github.com/opensbom-generator/spdx-sbom-generator/archive/refs/tags/v0.0.15.tar.gz"
	spdxsbomgenerator_cmd_tar := exec.Command("curl", "-L", spdxsbomgenerator_tar_url, "-o", "package.tar.gz")
	err := spdxsbomgenerator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spdxsbomgenerator_zip_url := "https://github.com/opensbom-generator/spdx-sbom-generator/archive/refs/tags/v0.0.15.zip"
	spdxsbomgenerator_cmd_zip := exec.Command("curl", "-L", spdxsbomgenerator_zip_url, "-o", "package.zip")
	err = spdxsbomgenerator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spdxsbomgenerator_bin_url := "https://github.com/opensbom-generator/spdx-sbom-generator/archive/refs/tags/v0.0.15.bin"
	spdxsbomgenerator_cmd_bin := exec.Command("curl", "-L", spdxsbomgenerator_bin_url, "-o", "binary.bin")
	err = spdxsbomgenerator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spdxsbomgenerator_src_url := "https://github.com/opensbom-generator/spdx-sbom-generator/archive/refs/tags/v0.0.15.src.tar.gz"
	spdxsbomgenerator_cmd_src := exec.Command("curl", "-L", spdxsbomgenerator_src_url, "-o", "source.tar.gz")
	err = spdxsbomgenerator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spdxsbomgenerator_cmd_direct := exec.Command("./binary")
	err = spdxsbomgenerator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
