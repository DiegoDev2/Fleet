package main

// Ufbt - Compact tool for building and debugging applications for Flipper Zero
// Homepage: https://pypi.org/project/ufbt/

import (
	"fmt"
	
	"os/exec"
)

func installUfbt() {
	// Método 1: Descargar y extraer .tar.gz
	ufbt_tar_url := "https://files.pythonhosted.org/packages/59/3b/013525f91836171870c49a53db8d2f772b5d32e682c0d25d0d0481c9bb51/ufbt-0.2.6.tar.gz"
	ufbt_cmd_tar := exec.Command("curl", "-L", ufbt_tar_url, "-o", "package.tar.gz")
	err := ufbt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ufbt_zip_url := "https://files.pythonhosted.org/packages/59/3b/013525f91836171870c49a53db8d2f772b5d32e682c0d25d0d0481c9bb51/ufbt-0.2.6.zip"
	ufbt_cmd_zip := exec.Command("curl", "-L", ufbt_zip_url, "-o", "package.zip")
	err = ufbt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ufbt_bin_url := "https://files.pythonhosted.org/packages/59/3b/013525f91836171870c49a53db8d2f772b5d32e682c0d25d0d0481c9bb51/ufbt-0.2.6.bin"
	ufbt_cmd_bin := exec.Command("curl", "-L", ufbt_bin_url, "-o", "binary.bin")
	err = ufbt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ufbt_src_url := "https://files.pythonhosted.org/packages/59/3b/013525f91836171870c49a53db8d2f772b5d32e682c0d25d0d0481c9bb51/ufbt-0.2.6.src.tar.gz"
	ufbt_cmd_src := exec.Command("curl", "-L", ufbt_src_url, "-o", "source.tar.gz")
	err = ufbt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ufbt_cmd_direct := exec.Command("./binary")
	err = ufbt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
