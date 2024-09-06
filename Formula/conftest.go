package main

// Conftest - Test your configuration files using Open Policy Agent
// Homepage: https://www.conftest.dev/

import (
	"fmt"
	
	"os/exec"
)

func installConftest() {
	// Método 1: Descargar y extraer .tar.gz
	conftest_tar_url := "https://github.com/open-policy-agent/conftest/archive/refs/tags/v0.55.0.tar.gz"
	conftest_cmd_tar := exec.Command("curl", "-L", conftest_tar_url, "-o", "package.tar.gz")
	err := conftest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	conftest_zip_url := "https://github.com/open-policy-agent/conftest/archive/refs/tags/v0.55.0.zip"
	conftest_cmd_zip := exec.Command("curl", "-L", conftest_zip_url, "-o", "package.zip")
	err = conftest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	conftest_bin_url := "https://github.com/open-policy-agent/conftest/archive/refs/tags/v0.55.0.bin"
	conftest_cmd_bin := exec.Command("curl", "-L", conftest_bin_url, "-o", "binary.bin")
	err = conftest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	conftest_src_url := "https://github.com/open-policy-agent/conftest/archive/refs/tags/v0.55.0.src.tar.gz"
	conftest_cmd_src := exec.Command("curl", "-L", conftest_src_url, "-o", "source.tar.gz")
	err = conftest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	conftest_cmd_direct := exec.Command("./binary")
	err = conftest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
