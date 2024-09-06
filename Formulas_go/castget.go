package main

// Castget - Command-line podcast and RSS enclosure downloader
// Homepage: https://castget.johndal.com/

import (
	"fmt"
	
	"os/exec"
)

func installCastget() {
	// Método 1: Descargar y extraer .tar.gz
	castget_tar_url := "https://download.savannah.gnu.org/releases/castget/castget-2.0.1.tar.bz2"
	castget_cmd_tar := exec.Command("curl", "-L", castget_tar_url, "-o", "package.tar.gz")
	err := castget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	castget_zip_url := "https://download.savannah.gnu.org/releases/castget/castget-2.0.1.tar.bz2"
	castget_cmd_zip := exec.Command("curl", "-L", castget_zip_url, "-o", "package.zip")
	err = castget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	castget_bin_url := "https://download.savannah.gnu.org/releases/castget/castget-2.0.1.tar.bz2"
	castget_cmd_bin := exec.Command("curl", "-L", castget_bin_url, "-o", "binary.bin")
	err = castget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	castget_src_url := "https://download.savannah.gnu.org/releases/castget/castget-2.0.1.tar.bz2"
	castget_cmd_src := exec.Command("curl", "-L", castget_src_url, "-o", "source.tar.gz")
	err = castget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	castget_cmd_direct := exec.Command("./binary")
	err = castget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: id3lib")
exec.Command("latte", "install", "id3lib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
