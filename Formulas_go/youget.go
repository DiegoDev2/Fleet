package main

// YouGet - Dumb downloader that scrapes the web
// Homepage: https://you-get.org/

import (
	"fmt"
	
	"os/exec"
)

func installYouGet() {
	// Método 1: Descargar y extraer .tar.gz
	youget_tar_url := "https://files.pythonhosted.org/packages/09/1e/96540e807ec3b103625e9660e7a2c7a7eb9accb1b90bf85156ff50e2dfd3/you_get-0.4.1718.tar.gz"
	youget_cmd_tar := exec.Command("curl", "-L", youget_tar_url, "-o", "package.tar.gz")
	err := youget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	youget_zip_url := "https://files.pythonhosted.org/packages/09/1e/96540e807ec3b103625e9660e7a2c7a7eb9accb1b90bf85156ff50e2dfd3/you_get-0.4.1718.zip"
	youget_cmd_zip := exec.Command("curl", "-L", youget_zip_url, "-o", "package.zip")
	err = youget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	youget_bin_url := "https://files.pythonhosted.org/packages/09/1e/96540e807ec3b103625e9660e7a2c7a7eb9accb1b90bf85156ff50e2dfd3/you_get-0.4.1718.bin"
	youget_cmd_bin := exec.Command("curl", "-L", youget_bin_url, "-o", "binary.bin")
	err = youget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	youget_src_url := "https://files.pythonhosted.org/packages/09/1e/96540e807ec3b103625e9660e7a2c7a7eb9accb1b90bf85156ff50e2dfd3/you_get-0.4.1718.src.tar.gz"
	youget_cmd_src := exec.Command("curl", "-L", youget_src_url, "-o", "source.tar.gz")
	err = youget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	youget_cmd_direct := exec.Command("./binary")
	err = youget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rtmpdump")
exec.Command("latte", "install", "rtmpdump")
}
