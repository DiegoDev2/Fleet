package main

// Pter - Your console and graphical UI to manage your todo.txt file(s)
// Homepage: https://vonshednob.cc/pter/

import (
	"fmt"
	
	"os/exec"
)

func installPter() {
	// Método 1: Descargar y extraer .tar.gz
	pter_tar_url := "https://files.pythonhosted.org/packages/d9/35/247e5568d1e500266bda2601df5b5169aec86bc421e76df298eeb2678fcf/pter-3.17.1.tar.gz"
	pter_cmd_tar := exec.Command("curl", "-L", pter_tar_url, "-o", "package.tar.gz")
	err := pter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pter_zip_url := "https://files.pythonhosted.org/packages/d9/35/247e5568d1e500266bda2601df5b5169aec86bc421e76df298eeb2678fcf/pter-3.17.1.zip"
	pter_cmd_zip := exec.Command("curl", "-L", pter_zip_url, "-o", "package.zip")
	err = pter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pter_bin_url := "https://files.pythonhosted.org/packages/d9/35/247e5568d1e500266bda2601df5b5169aec86bc421e76df298eeb2678fcf/pter-3.17.1.bin"
	pter_cmd_bin := exec.Command("curl", "-L", pter_bin_url, "-o", "binary.bin")
	err = pter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pter_src_url := "https://files.pythonhosted.org/packages/d9/35/247e5568d1e500266bda2601df5b5169aec86bc421e76df298eeb2678fcf/pter-3.17.1.src.tar.gz"
	pter_cmd_src := exec.Command("curl", "-L", pter_src_url, "-o", "source.tar.gz")
	err = pter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pter_cmd_direct := exec.Command("./binary")
	err = pter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
