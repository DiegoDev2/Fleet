package main

// Woob - Web Outside of Browsers
// Homepage: https://woob.tech/

import (
	"fmt"
	
	"os/exec"
)

func installWoob() {
	// Método 1: Descargar y extraer .tar.gz
	woob_tar_url := "https://files.pythonhosted.org/packages/cf/10/3eb104a43ab4ff3109109883382bdfee663412e8fda2967d0ab220479240/woob-3.6.tar.gz"
	woob_cmd_tar := exec.Command("curl", "-L", woob_tar_url, "-o", "package.tar.gz")
	err := woob_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	woob_zip_url := "https://files.pythonhosted.org/packages/cf/10/3eb104a43ab4ff3109109883382bdfee663412e8fda2967d0ab220479240/woob-3.6.zip"
	woob_cmd_zip := exec.Command("curl", "-L", woob_zip_url, "-o", "package.zip")
	err = woob_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	woob_bin_url := "https://files.pythonhosted.org/packages/cf/10/3eb104a43ab4ff3109109883382bdfee663412e8fda2967d0ab220479240/woob-3.6.bin"
	woob_cmd_bin := exec.Command("curl", "-L", woob_bin_url, "-o", "binary.bin")
	err = woob_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	woob_src_url := "https://files.pythonhosted.org/packages/cf/10/3eb104a43ab4ff3109109883382bdfee663412e8fda2967d0ab220479240/woob-3.6.src.tar.gz"
	woob_cmd_src := exec.Command("curl", "-L", woob_src_url, "-o", "source.tar.gz")
	err = woob_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	woob_cmd_direct := exec.Command("./binary")
	err = woob_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: gnupg")
exec.Command("latte", "install", "gnupg")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
