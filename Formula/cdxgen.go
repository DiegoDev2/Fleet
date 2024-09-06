package main

// Cdxgen - Creates CycloneDX Software Bill-of-Materials (SBOM) for projects
// Homepage: https://github.com/CycloneDX/cdxgen

import (
	"fmt"
	
	"os/exec"
)

func installCdxgen() {
	// Método 1: Descargar y extraer .tar.gz
	cdxgen_tar_url := "https://registry.npmjs.org/@cyclonedx/cdxgen/-/cdxgen-10.9.8.tgz"
	cdxgen_cmd_tar := exec.Command("curl", "-L", cdxgen_tar_url, "-o", "package.tar.gz")
	err := cdxgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdxgen_zip_url := "https://registry.npmjs.org/@cyclonedx/cdxgen/-/cdxgen-10.9.8.tgz"
	cdxgen_cmd_zip := exec.Command("curl", "-L", cdxgen_zip_url, "-o", "package.zip")
	err = cdxgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdxgen_bin_url := "https://registry.npmjs.org/@cyclonedx/cdxgen/-/cdxgen-10.9.8.tgz"
	cdxgen_cmd_bin := exec.Command("curl", "-L", cdxgen_bin_url, "-o", "binary.bin")
	err = cdxgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdxgen_src_url := "https://registry.npmjs.org/@cyclonedx/cdxgen/-/cdxgen-10.9.8.tgz"
	cdxgen_cmd_src := exec.Command("curl", "-L", cdxgen_src_url, "-o", "source.tar.gz")
	err = cdxgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdxgen_cmd_direct := exec.Command("./binary")
	err = cdxgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
