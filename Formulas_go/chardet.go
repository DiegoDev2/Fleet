package main

// Chardet - Python character encoding detector
// Homepage: https://chardet.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installChardet() {
	// Método 1: Descargar y extraer .tar.gz
	chardet_tar_url := "https://files.pythonhosted.org/packages/f3/0d/f7b6ab21ec75897ed80c17d79b15951a719226b9fababf1e40ea74d69079/chardet-5.2.0.tar.gz"
	chardet_cmd_tar := exec.Command("curl", "-L", chardet_tar_url, "-o", "package.tar.gz")
	err := chardet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chardet_zip_url := "https://files.pythonhosted.org/packages/f3/0d/f7b6ab21ec75897ed80c17d79b15951a719226b9fababf1e40ea74d69079/chardet-5.2.0.zip"
	chardet_cmd_zip := exec.Command("curl", "-L", chardet_zip_url, "-o", "package.zip")
	err = chardet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chardet_bin_url := "https://files.pythonhosted.org/packages/f3/0d/f7b6ab21ec75897ed80c17d79b15951a719226b9fababf1e40ea74d69079/chardet-5.2.0.bin"
	chardet_cmd_bin := exec.Command("curl", "-L", chardet_bin_url, "-o", "binary.bin")
	err = chardet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chardet_src_url := "https://files.pythonhosted.org/packages/f3/0d/f7b6ab21ec75897ed80c17d79b15951a719226b9fababf1e40ea74d69079/chardet-5.2.0.src.tar.gz"
	chardet_cmd_src := exec.Command("curl", "-L", chardet_src_url, "-o", "source.tar.gz")
	err = chardet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chardet_cmd_direct := exec.Command("./binary")
	err = chardet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
