package main

// Vint - Vim script Language Lint
// Homepage: https://github.com/Vimjas/vint

import (
	"fmt"
	
	"os/exec"
)

func installVint() {
	// Método 1: Descargar y extraer .tar.gz
	vint_tar_url := "https://files.pythonhosted.org/packages/9c/c7/d5fbe5f778edee83cba3aea8cc3308db327e4c161e0656e861b9cc2cb859/vim-vint-0.3.21.tar.gz"
	vint_cmd_tar := exec.Command("curl", "-L", vint_tar_url, "-o", "package.tar.gz")
	err := vint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vint_zip_url := "https://files.pythonhosted.org/packages/9c/c7/d5fbe5f778edee83cba3aea8cc3308db327e4c161e0656e861b9cc2cb859/vim-vint-0.3.21.zip"
	vint_cmd_zip := exec.Command("curl", "-L", vint_zip_url, "-o", "package.zip")
	err = vint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vint_bin_url := "https://files.pythonhosted.org/packages/9c/c7/d5fbe5f778edee83cba3aea8cc3308db327e4c161e0656e861b9cc2cb859/vim-vint-0.3.21.bin"
	vint_cmd_bin := exec.Command("curl", "-L", vint_bin_url, "-o", "binary.bin")
	err = vint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vint_src_url := "https://files.pythonhosted.org/packages/9c/c7/d5fbe5f778edee83cba3aea8cc3308db327e4c161e0656e861b9cc2cb859/vim-vint-0.3.21.src.tar.gz"
	vint_cmd_src := exec.Command("curl", "-L", vint_src_url, "-o", "source.tar.gz")
	err = vint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vint_cmd_direct := exec.Command("./binary")
	err = vint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
