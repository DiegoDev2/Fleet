package main

// Unoconv - Convert between any document format supported by OpenOffice
// Homepage: https://github.com/unoconv/unoconv

import (
	"fmt"
	
	"os/exec"
)

func installUnoconv() {
	// Método 1: Descargar y extraer .tar.gz
	unoconv_tar_url := "https://files.pythonhosted.org/packages/ab/40/b4cab1140087f3f07b2f6d7cb9ca1c14b9bdbb525d2d83a3b29c924fe9ae/unoconv-0.9.0.tar.gz"
	unoconv_cmd_tar := exec.Command("curl", "-L", unoconv_tar_url, "-o", "package.tar.gz")
	err := unoconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unoconv_zip_url := "https://files.pythonhosted.org/packages/ab/40/b4cab1140087f3f07b2f6d7cb9ca1c14b9bdbb525d2d83a3b29c924fe9ae/unoconv-0.9.0.zip"
	unoconv_cmd_zip := exec.Command("curl", "-L", unoconv_zip_url, "-o", "package.zip")
	err = unoconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unoconv_bin_url := "https://files.pythonhosted.org/packages/ab/40/b4cab1140087f3f07b2f6d7cb9ca1c14b9bdbb525d2d83a3b29c924fe9ae/unoconv-0.9.0.bin"
	unoconv_cmd_bin := exec.Command("curl", "-L", unoconv_bin_url, "-o", "binary.bin")
	err = unoconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unoconv_src_url := "https://files.pythonhosted.org/packages/ab/40/b4cab1140087f3f07b2f6d7cb9ca1c14b9bdbb525d2d83a3b29c924fe9ae/unoconv-0.9.0.src.tar.gz"
	unoconv_cmd_src := exec.Command("curl", "-L", unoconv_src_url, "-o", "source.tar.gz")
	err = unoconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unoconv_cmd_direct := exec.Command("./binary")
	err = unoconv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
