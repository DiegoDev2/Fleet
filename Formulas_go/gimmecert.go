package main

// Gimmecert - Quickly issue X.509 server and client certificates using locally-generated CA
// Homepage: https://projects.majic.rs/gimmecert

import (
	"fmt"
	
	"os/exec"
)

func installGimmecert() {
	// Método 1: Descargar y extraer .tar.gz
	gimmecert_tar_url := "https://files.pythonhosted.org/packages/94/b3/f8d0d4fc8951d7ff02f1d3653ba446ad0edf14ab1a18cff4fbe1d1b62086/gimmecert-1.0.0.tar.gz"
	gimmecert_cmd_tar := exec.Command("curl", "-L", gimmecert_tar_url, "-o", "package.tar.gz")
	err := gimmecert_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gimmecert_zip_url := "https://files.pythonhosted.org/packages/94/b3/f8d0d4fc8951d7ff02f1d3653ba446ad0edf14ab1a18cff4fbe1d1b62086/gimmecert-1.0.0.zip"
	gimmecert_cmd_zip := exec.Command("curl", "-L", gimmecert_zip_url, "-o", "package.zip")
	err = gimmecert_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gimmecert_bin_url := "https://files.pythonhosted.org/packages/94/b3/f8d0d4fc8951d7ff02f1d3653ba446ad0edf14ab1a18cff4fbe1d1b62086/gimmecert-1.0.0.bin"
	gimmecert_cmd_bin := exec.Command("curl", "-L", gimmecert_bin_url, "-o", "binary.bin")
	err = gimmecert_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gimmecert_src_url := "https://files.pythonhosted.org/packages/94/b3/f8d0d4fc8951d7ff02f1d3653ba446ad0edf14ab1a18cff4fbe1d1b62086/gimmecert-1.0.0.src.tar.gz"
	gimmecert_cmd_src := exec.Command("curl", "-L", gimmecert_src_url, "-o", "source.tar.gz")
	err = gimmecert_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gimmecert_cmd_direct := exec.Command("./binary")
	err = gimmecert_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
