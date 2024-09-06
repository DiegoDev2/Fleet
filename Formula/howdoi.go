package main

// Howdoi - Instant coding answers via the command-line
// Homepage: https://github.com/gleitz/howdoi

import (
	"fmt"
	
	"os/exec"
)

func installHowdoi() {
	// Método 1: Descargar y extraer .tar.gz
	howdoi_tar_url := "https://files.pythonhosted.org/packages/6d/43/0e8166583575bd500c0f8f1a4ab9429af9466feb6fcdc006e88de8fd23e9/howdoi-2.0.20.tar.gz"
	howdoi_cmd_tar := exec.Command("curl", "-L", howdoi_tar_url, "-o", "package.tar.gz")
	err := howdoi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	howdoi_zip_url := "https://files.pythonhosted.org/packages/6d/43/0e8166583575bd500c0f8f1a4ab9429af9466feb6fcdc006e88de8fd23e9/howdoi-2.0.20.zip"
	howdoi_cmd_zip := exec.Command("curl", "-L", howdoi_zip_url, "-o", "package.zip")
	err = howdoi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	howdoi_bin_url := "https://files.pythonhosted.org/packages/6d/43/0e8166583575bd500c0f8f1a4ab9429af9466feb6fcdc006e88de8fd23e9/howdoi-2.0.20.bin"
	howdoi_cmd_bin := exec.Command("curl", "-L", howdoi_bin_url, "-o", "binary.bin")
	err = howdoi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	howdoi_src_url := "https://files.pythonhosted.org/packages/6d/43/0e8166583575bd500c0f8f1a4ab9429af9466feb6fcdc006e88de8fd23e9/howdoi-2.0.20.src.tar.gz"
	howdoi_cmd_src := exec.Command("curl", "-L", howdoi_src_url, "-o", "source.tar.gz")
	err = howdoi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	howdoi_cmd_direct := exec.Command("./binary")
	err = howdoi_cmd_direct.Run()
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
