package main

// Pcb2gcode - Command-line tool for isolation, routing and drilling of PCBs
// Homepage: https://github.com/pcb2gcode/pcb2gcode

import (
	"fmt"
	
	"os/exec"
)

func installPcb2gcode() {
	// Método 1: Descargar y extraer .tar.gz
	pcb2gcode_tar_url := "https://github.com/pcb2gcode/pcb2gcode/archive/refs/tags/v2.5.0.tar.gz"
	pcb2gcode_cmd_tar := exec.Command("curl", "-L", pcb2gcode_tar_url, "-o", "package.tar.gz")
	err := pcb2gcode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcb2gcode_zip_url := "https://github.com/pcb2gcode/pcb2gcode/archive/refs/tags/v2.5.0.zip"
	pcb2gcode_cmd_zip := exec.Command("curl", "-L", pcb2gcode_zip_url, "-o", "package.zip")
	err = pcb2gcode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcb2gcode_bin_url := "https://github.com/pcb2gcode/pcb2gcode/archive/refs/tags/v2.5.0.bin"
	pcb2gcode_cmd_bin := exec.Command("curl", "-L", pcb2gcode_bin_url, "-o", "binary.bin")
	err = pcb2gcode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcb2gcode_src_url := "https://github.com/pcb2gcode/pcb2gcode/archive/refs/tags/v2.5.0.src.tar.gz"
	pcb2gcode_cmd_src := exec.Command("curl", "-L", pcb2gcode_src_url, "-o", "source.tar.gz")
	err = pcb2gcode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcb2gcode_cmd_direct := exec.Command("./binary")
	err = pcb2gcode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cairomm@1.14")
	exec.Command("latte", "install", "cairomm@1.14").Run()
	fmt.Println("Instalando dependencia: glibmm@2.66")
	exec.Command("latte", "install", "glibmm@2.66").Run()
	fmt.Println("Instalando dependencia: gtkmm")
	exec.Command("latte", "install", "gtkmm").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pangomm@2.46")
	exec.Command("latte", "install", "pangomm@2.46").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gerbv")
	exec.Command("latte", "install", "gerbv").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
