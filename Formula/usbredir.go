package main

// Usbredir - USB traffic redirection library
// Homepage: https://www.spice-space.org

import (
	"fmt"
	
	"os/exec"
)

func installUsbredir() {
	// Método 1: Descargar y extraer .tar.gz
	usbredir_tar_url := "https://www.spice-space.org/download/usbredir/usbredir-0.14.0.tar.xz"
	usbredir_cmd_tar := exec.Command("curl", "-L", usbredir_tar_url, "-o", "package.tar.gz")
	err := usbredir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	usbredir_zip_url := "https://www.spice-space.org/download/usbredir/usbredir-0.14.0.tar.xz"
	usbredir_cmd_zip := exec.Command("curl", "-L", usbredir_zip_url, "-o", "package.zip")
	err = usbredir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	usbredir_bin_url := "https://www.spice-space.org/download/usbredir/usbredir-0.14.0.tar.xz"
	usbredir_cmd_bin := exec.Command("curl", "-L", usbredir_bin_url, "-o", "binary.bin")
	err = usbredir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	usbredir_src_url := "https://www.spice-space.org/download/usbredir/usbredir-0.14.0.tar.xz"
	usbredir_cmd_src := exec.Command("curl", "-L", usbredir_src_url, "-o", "source.tar.gz")
	err = usbredir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	usbredir_cmd_direct := exec.Command("./binary")
	err = usbredir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
