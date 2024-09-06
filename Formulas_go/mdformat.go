package main

// Mdformat - CommonMark compliant Markdown formatter
// Homepage: https://mdformat.readthedocs.io/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installMdformat() {
	// Método 1: Descargar y extraer .tar.gz
	mdformat_tar_url := "https://files.pythonhosted.org/packages/df/86/6374cc48a89862cfc8e350a65d6af47792e83e7684f13e1222afce110a41/mdformat-0.7.17.tar.gz"
	mdformat_cmd_tar := exec.Command("curl", "-L", mdformat_tar_url, "-o", "package.tar.gz")
	err := mdformat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdformat_zip_url := "https://files.pythonhosted.org/packages/df/86/6374cc48a89862cfc8e350a65d6af47792e83e7684f13e1222afce110a41/mdformat-0.7.17.zip"
	mdformat_cmd_zip := exec.Command("curl", "-L", mdformat_zip_url, "-o", "package.zip")
	err = mdformat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdformat_bin_url := "https://files.pythonhosted.org/packages/df/86/6374cc48a89862cfc8e350a65d6af47792e83e7684f13e1222afce110a41/mdformat-0.7.17.bin"
	mdformat_cmd_bin := exec.Command("curl", "-L", mdformat_bin_url, "-o", "binary.bin")
	err = mdformat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdformat_src_url := "https://files.pythonhosted.org/packages/df/86/6374cc48a89862cfc8e350a65d6af47792e83e7684f13e1222afce110a41/mdformat-0.7.17.src.tar.gz"
	mdformat_cmd_src := exec.Command("curl", "-L", mdformat_src_url, "-o", "source.tar.gz")
	err = mdformat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdformat_cmd_direct := exec.Command("./binary")
	err = mdformat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
