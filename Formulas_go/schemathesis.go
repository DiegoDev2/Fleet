package main

// Schemathesis - Testing tool for web applications with specs
// Homepage: https://schemathesis.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installSchemathesis() {
	// Método 1: Descargar y extraer .tar.gz
	schemathesis_tar_url := "https://files.pythonhosted.org/packages/b6/fd/126fad05418364b10b8ab78713a4adcba4ea55a8d77be4bbfec5f206efef/schemathesis-3.35.4.tar.gz"
	schemathesis_cmd_tar := exec.Command("curl", "-L", schemathesis_tar_url, "-o", "package.tar.gz")
	err := schemathesis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	schemathesis_zip_url := "https://files.pythonhosted.org/packages/b6/fd/126fad05418364b10b8ab78713a4adcba4ea55a8d77be4bbfec5f206efef/schemathesis-3.35.4.zip"
	schemathesis_cmd_zip := exec.Command("curl", "-L", schemathesis_zip_url, "-o", "package.zip")
	err = schemathesis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	schemathesis_bin_url := "https://files.pythonhosted.org/packages/b6/fd/126fad05418364b10b8ab78713a4adcba4ea55a8d77be4bbfec5f206efef/schemathesis-3.35.4.bin"
	schemathesis_cmd_bin := exec.Command("curl", "-L", schemathesis_bin_url, "-o", "binary.bin")
	err = schemathesis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	schemathesis_src_url := "https://files.pythonhosted.org/packages/b6/fd/126fad05418364b10b8ab78713a4adcba4ea55a8d77be4bbfec5f206efef/schemathesis-3.35.4.src.tar.gz"
	schemathesis_cmd_src := exec.Command("curl", "-L", schemathesis_src_url, "-o", "source.tar.gz")
	err = schemathesis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	schemathesis_cmd_direct := exec.Command("./binary")
	err = schemathesis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
