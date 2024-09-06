package main

// Sickchill - Automatic Video Library Manager for TV Shows
// Homepage: https://sickchill.github.io

import (
	"fmt"
	
	"os/exec"
)

func installSickchill() {
	// Método 1: Descargar y extraer .tar.gz
	sickchill_tar_url := "https://files.pythonhosted.org/packages/31/fc/337b2989dc67bbb505cea34a05c029cbba3056311177586835f704ddc13a/sickchill-2024.3.1.tar.gz"
	sickchill_cmd_tar := exec.Command("curl", "-L", sickchill_tar_url, "-o", "package.tar.gz")
	err := sickchill_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sickchill_zip_url := "https://files.pythonhosted.org/packages/31/fc/337b2989dc67bbb505cea34a05c029cbba3056311177586835f704ddc13a/sickchill-2024.3.1.zip"
	sickchill_cmd_zip := exec.Command("curl", "-L", sickchill_zip_url, "-o", "package.zip")
	err = sickchill_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sickchill_bin_url := "https://files.pythonhosted.org/packages/31/fc/337b2989dc67bbb505cea34a05c029cbba3056311177586835f704ddc13a/sickchill-2024.3.1.bin"
	sickchill_cmd_bin := exec.Command("curl", "-L", sickchill_bin_url, "-o", "binary.bin")
	err = sickchill_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sickchill_src_url := "https://files.pythonhosted.org/packages/31/fc/337b2989dc67bbb505cea34a05c029cbba3056311177586835f704ddc13a/sickchill-2024.3.1.src.tar.gz"
	sickchill_cmd_src := exec.Command("curl", "-L", sickchill_src_url, "-o", "source.tar.gz")
	err = sickchill_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sickchill_cmd_direct := exec.Command("./binary")
	err = sickchill_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
