package main

// Kaskade - TUI for Kafka
// Homepage: https://github.com/sauljabin/kaskade

import (
	"fmt"
	
	"os/exec"
)

func installKaskade() {
	// Método 1: Descargar y extraer .tar.gz
	kaskade_tar_url := "https://files.pythonhosted.org/packages/67/3f/d02b5505f16ec84280da9ced23529a16ea6a1565add37810b9f686adcd6b/kaskade-2.3.5.tar.gz"
	kaskade_cmd_tar := exec.Command("curl", "-L", kaskade_tar_url, "-o", "package.tar.gz")
	err := kaskade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kaskade_zip_url := "https://files.pythonhosted.org/packages/67/3f/d02b5505f16ec84280da9ced23529a16ea6a1565add37810b9f686adcd6b/kaskade-2.3.5.zip"
	kaskade_cmd_zip := exec.Command("curl", "-L", kaskade_zip_url, "-o", "package.zip")
	err = kaskade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kaskade_bin_url := "https://files.pythonhosted.org/packages/67/3f/d02b5505f16ec84280da9ced23529a16ea6a1565add37810b9f686adcd6b/kaskade-2.3.5.bin"
	kaskade_cmd_bin := exec.Command("curl", "-L", kaskade_bin_url, "-o", "binary.bin")
	err = kaskade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kaskade_src_url := "https://files.pythonhosted.org/packages/67/3f/d02b5505f16ec84280da9ced23529a16ea6a1565add37810b9f686adcd6b/kaskade-2.3.5.src.tar.gz"
	kaskade_cmd_src := exec.Command("curl", "-L", kaskade_src_url, "-o", "source.tar.gz")
	err = kaskade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kaskade_cmd_direct := exec.Command("./binary")
	err = kaskade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: librdkafka")
exec.Command("latte", "install", "librdkafka")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
