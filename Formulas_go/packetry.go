package main

// Packetry - Fast, intuitive USB 2.0 protocol analysis application for use with Cynthion
// Homepage: https://github.com/greatscottgadgets/packetry

import (
	"fmt"
	
	"os/exec"
)

func installPacketry() {
	// Método 1: Descargar y extraer .tar.gz
	packetry_tar_url := "https://github.com/greatscottgadgets/packetry/archive/refs/tags/v0.2.2.tar.gz"
	packetry_cmd_tar := exec.Command("curl", "-L", packetry_tar_url, "-o", "package.tar.gz")
	err := packetry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	packetry_zip_url := "https://github.com/greatscottgadgets/packetry/archive/refs/tags/v0.2.2.zip"
	packetry_cmd_zip := exec.Command("curl", "-L", packetry_zip_url, "-o", "package.zip")
	err = packetry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	packetry_bin_url := "https://github.com/greatscottgadgets/packetry/archive/refs/tags/v0.2.2.bin"
	packetry_cmd_bin := exec.Command("curl", "-L", packetry_bin_url, "-o", "binary.bin")
	err = packetry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	packetry_src_url := "https://github.com/greatscottgadgets/packetry/archive/refs/tags/v0.2.2.src.tar.gz"
	packetry_cmd_src := exec.Command("curl", "-L", packetry_src_url, "-o", "source.tar.gz")
	err = packetry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	packetry_cmd_direct := exec.Command("./binary")
	err = packetry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: graphene")
exec.Command("latte", "install", "graphene")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
