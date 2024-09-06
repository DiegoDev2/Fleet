package main

// Pcb - Interactive printed circuit board editor
// Homepage: http://pcb.geda-project.org/

import (
	"fmt"
	
	"os/exec"
)

func installPcb() {
	// Método 1: Descargar y extraer .tar.gz
	pcb_tar_url := "https://downloads.sourceforge.net/project/pcb/pcb/pcb-4.3.0/pcb-4.3.0.tar.gz"
	pcb_cmd_tar := exec.Command("curl", "-L", pcb_tar_url, "-o", "package.tar.gz")
	err := pcb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcb_zip_url := "https://downloads.sourceforge.net/project/pcb/pcb/pcb-4.3.0/pcb-4.3.0.zip"
	pcb_cmd_zip := exec.Command("curl", "-L", pcb_zip_url, "-o", "package.zip")
	err = pcb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcb_bin_url := "https://downloads.sourceforge.net/project/pcb/pcb/pcb-4.3.0/pcb-4.3.0.bin"
	pcb_cmd_bin := exec.Command("curl", "-L", pcb_bin_url, "-o", "binary.bin")
	err = pcb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcb_src_url := "https://downloads.sourceforge.net/project/pcb/pcb/pcb-4.3.0/pcb-4.3.0.src.tar.gz"
	pcb_cmd_src := exec.Command("curl", "-L", pcb_src_url, "-o", "source.tar.gz")
	err = pcb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcb_cmd_direct := exec.Command("./binary")
	err = pcb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: gtkglext")
exec.Command("latte", "install", "gtkglext")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
