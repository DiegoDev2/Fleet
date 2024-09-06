package main

// Arjun - HTTP parameter discovery suite
// Homepage: https://github.com/s0md3v/Arjun

import (
	"fmt"
	
	"os/exec"
)

func installArjun() {
	// Método 1: Descargar y extraer .tar.gz
	arjun_tar_url := "https://files.pythonhosted.org/packages/bb/97/ed0189286d98aaf92322a06e23b10fc6c298e0ee9a43cd69ab614a1f76cf/arjun-2.2.6.tar.gz"
	arjun_cmd_tar := exec.Command("curl", "-L", arjun_tar_url, "-o", "package.tar.gz")
	err := arjun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arjun_zip_url := "https://files.pythonhosted.org/packages/bb/97/ed0189286d98aaf92322a06e23b10fc6c298e0ee9a43cd69ab614a1f76cf/arjun-2.2.6.zip"
	arjun_cmd_zip := exec.Command("curl", "-L", arjun_zip_url, "-o", "package.zip")
	err = arjun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arjun_bin_url := "https://files.pythonhosted.org/packages/bb/97/ed0189286d98aaf92322a06e23b10fc6c298e0ee9a43cd69ab614a1f76cf/arjun-2.2.6.bin"
	arjun_cmd_bin := exec.Command("curl", "-L", arjun_bin_url, "-o", "binary.bin")
	err = arjun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arjun_src_url := "https://files.pythonhosted.org/packages/bb/97/ed0189286d98aaf92322a06e23b10fc6c298e0ee9a43cd69ab614a1f76cf/arjun-2.2.6.src.tar.gz"
	arjun_cmd_src := exec.Command("curl", "-L", arjun_src_url, "-o", "source.tar.gz")
	err = arjun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arjun_cmd_direct := exec.Command("./binary")
	err = arjun_cmd_direct.Run()
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
