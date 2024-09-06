package main

// Dstack - ML workflow orchestration system designed for reproducibility and collaboration
// Homepage: https://dstack.ai/

import (
	"fmt"
	
	"os/exec"
)

func installDstack() {
	// Método 1: Descargar y extraer .tar.gz
	dstack_tar_url := "https://files.pythonhosted.org/packages/c4/d3/0fac0889db75f930d9bf32c2ae7da4bf73093446d7e2b5b719ef9e11e2e8/dstack-0.18.12.tar.gz"
	dstack_cmd_tar := exec.Command("curl", "-L", dstack_tar_url, "-o", "package.tar.gz")
	err := dstack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dstack_zip_url := "https://files.pythonhosted.org/packages/c4/d3/0fac0889db75f930d9bf32c2ae7da4bf73093446d7e2b5b719ef9e11e2e8/dstack-0.18.12.zip"
	dstack_cmd_zip := exec.Command("curl", "-L", dstack_zip_url, "-o", "package.zip")
	err = dstack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dstack_bin_url := "https://files.pythonhosted.org/packages/c4/d3/0fac0889db75f930d9bf32c2ae7da4bf73093446d7e2b5b719ef9e11e2e8/dstack-0.18.12.bin"
	dstack_cmd_bin := exec.Command("curl", "-L", dstack_bin_url, "-o", "binary.bin")
	err = dstack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dstack_src_url := "https://files.pythonhosted.org/packages/c4/d3/0fac0889db75f930d9bf32c2ae7da4bf73093446d7e2b5b719ef9e11e2e8/dstack-0.18.12.src.tar.gz"
	dstack_cmd_src := exec.Command("curl", "-L", dstack_src_url, "-o", "source.tar.gz")
	err = dstack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dstack_cmd_direct := exec.Command("./binary")
	err = dstack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
