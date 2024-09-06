package main

// Dxpy - DNAnexus toolkit utilities and platform API bindings for Python
// Homepage: https://github.com/dnanexus/dx-toolkit

import (
	"fmt"
	
	"os/exec"
)

func installDxpy() {
	// Método 1: Descargar y extraer .tar.gz
	dxpy_tar_url := "https://files.pythonhosted.org/packages/10/14/eaaa19a8970127db18dfd3c091cd995a5e2adeac2dadcc5d343dceba6867/dxpy-0.381.0.tar.gz"
	dxpy_cmd_tar := exec.Command("curl", "-L", dxpy_tar_url, "-o", "package.tar.gz")
	err := dxpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dxpy_zip_url := "https://files.pythonhosted.org/packages/10/14/eaaa19a8970127db18dfd3c091cd995a5e2adeac2dadcc5d343dceba6867/dxpy-0.381.0.zip"
	dxpy_cmd_zip := exec.Command("curl", "-L", dxpy_zip_url, "-o", "package.zip")
	err = dxpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dxpy_bin_url := "https://files.pythonhosted.org/packages/10/14/eaaa19a8970127db18dfd3c091cd995a5e2adeac2dadcc5d343dceba6867/dxpy-0.381.0.bin"
	dxpy_cmd_bin := exec.Command("curl", "-L", dxpy_bin_url, "-o", "binary.bin")
	err = dxpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dxpy_src_url := "https://files.pythonhosted.org/packages/10/14/eaaa19a8970127db18dfd3c091cd995a5e2adeac2dadcc5d343dceba6867/dxpy-0.381.0.src.tar.gz"
	dxpy_cmd_src := exec.Command("curl", "-L", dxpy_src_url, "-o", "source.tar.gz")
	err = dxpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dxpy_cmd_direct := exec.Command("./binary")
	err = dxpy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
