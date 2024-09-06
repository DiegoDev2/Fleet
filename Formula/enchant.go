package main

// Enchant - Spellchecker wrapping library
// Homepage: https://abiword.github.io/enchant/

import (
	"fmt"
	
	"os/exec"
)

func installEnchant() {
	// Método 1: Descargar y extraer .tar.gz
	enchant_tar_url := "https://github.com/AbiWord/enchant/releases/download/v2.8.2/enchant-2.8.2.tar.gz"
	enchant_cmd_tar := exec.Command("curl", "-L", enchant_tar_url, "-o", "package.tar.gz")
	err := enchant_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enchant_zip_url := "https://github.com/AbiWord/enchant/releases/download/v2.8.2/enchant-2.8.2.zip"
	enchant_cmd_zip := exec.Command("curl", "-L", enchant_zip_url, "-o", "package.zip")
	err = enchant_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enchant_bin_url := "https://github.com/AbiWord/enchant/releases/download/v2.8.2/enchant-2.8.2.bin"
	enchant_cmd_bin := exec.Command("curl", "-L", enchant_bin_url, "-o", "binary.bin")
	err = enchant_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enchant_src_url := "https://github.com/AbiWord/enchant/releases/download/v2.8.2/enchant-2.8.2.src.tar.gz"
	enchant_cmd_src := exec.Command("curl", "-L", enchant_src_url, "-o", "source.tar.gz")
	err = enchant_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enchant_cmd_direct := exec.Command("./binary")
	err = enchant_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: aspell")
	exec.Command("latte", "install", "aspell").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: groff")
	exec.Command("latte", "install", "groff").Run()
}
