package main

// NicotinePlus - Graphical client for the Soulseek peer-to-peer network
// Homepage: https://nicotine-plus.org

import (
	"fmt"
	
	"os/exec"
)

func installNicotinePlus() {
	// Método 1: Descargar y extraer .tar.gz
	nicotineplus_tar_url := "https://files.pythonhosted.org/packages/63/1c/73f765da20b5b7e3579f6099490a9c4ac93e7c6341f97cf51d53ea0df49f/nicotine_plus-3.3.4.tar.gz"
	nicotineplus_cmd_tar := exec.Command("curl", "-L", nicotineplus_tar_url, "-o", "package.tar.gz")
	err := nicotineplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nicotineplus_zip_url := "https://files.pythonhosted.org/packages/63/1c/73f765da20b5b7e3579f6099490a9c4ac93e7c6341f97cf51d53ea0df49f/nicotine_plus-3.3.4.zip"
	nicotineplus_cmd_zip := exec.Command("curl", "-L", nicotineplus_zip_url, "-o", "package.zip")
	err = nicotineplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nicotineplus_bin_url := "https://files.pythonhosted.org/packages/63/1c/73f765da20b5b7e3579f6099490a9c4ac93e7c6341f97cf51d53ea0df49f/nicotine_plus-3.3.4.bin"
	nicotineplus_cmd_bin := exec.Command("curl", "-L", nicotineplus_bin_url, "-o", "binary.bin")
	err = nicotineplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nicotineplus_src_url := "https://files.pythonhosted.org/packages/63/1c/73f765da20b5b7e3579f6099490a9c4ac93e7c6341f97cf51d53ea0df49f/nicotine_plus-3.3.4.src.tar.gz"
	nicotineplus_cmd_src := exec.Command("curl", "-L", nicotineplus_src_url, "-o", "source.tar.gz")
	err = nicotineplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nicotineplus_cmd_direct := exec.Command("./binary")
	err = nicotineplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: libadwaita")
	exec.Command("latte", "install", "libadwaita").Run()
	fmt.Println("Instalando dependencia: py3cairo")
	exec.Command("latte", "install", "py3cairo").Run()
	fmt.Println("Instalando dependencia: pygobject3")
	exec.Command("latte", "install", "pygobject3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
