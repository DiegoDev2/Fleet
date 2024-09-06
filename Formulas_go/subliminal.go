package main

// Subliminal - Library to search and download subtitles
// Homepage: https://subliminal.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installSubliminal() {
	// Método 1: Descargar y extraer .tar.gz
	subliminal_tar_url := "https://files.pythonhosted.org/packages/e3/24/36cdb82e90afc602e2ed36c34e022ca545d35f5be9aa7ef9ddb0af7967b2/subliminal-2.2.1.tar.gz"
	subliminal_cmd_tar := exec.Command("curl", "-L", subliminal_tar_url, "-o", "package.tar.gz")
	err := subliminal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	subliminal_zip_url := "https://files.pythonhosted.org/packages/e3/24/36cdb82e90afc602e2ed36c34e022ca545d35f5be9aa7ef9ddb0af7967b2/subliminal-2.2.1.zip"
	subliminal_cmd_zip := exec.Command("curl", "-L", subliminal_zip_url, "-o", "package.zip")
	err = subliminal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	subliminal_bin_url := "https://files.pythonhosted.org/packages/e3/24/36cdb82e90afc602e2ed36c34e022ca545d35f5be9aa7ef9ddb0af7967b2/subliminal-2.2.1.bin"
	subliminal_cmd_bin := exec.Command("curl", "-L", subliminal_bin_url, "-o", "binary.bin")
	err = subliminal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	subliminal_src_url := "https://files.pythonhosted.org/packages/e3/24/36cdb82e90afc602e2ed36c34e022ca545d35f5be9aa7ef9ddb0af7967b2/subliminal-2.2.1.src.tar.gz"
	subliminal_cmd_src := exec.Command("curl", "-L", subliminal_src_url, "-o", "source.tar.gz")
	err = subliminal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	subliminal_cmd_direct := exec.Command("./binary")
	err = subliminal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
