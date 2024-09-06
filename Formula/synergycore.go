package main

// SynergyCore - Synergy, the keyboard and mouse sharing tool
// Homepage: https://symless.com/synergy

import (
	"fmt"
	
	"os/exec"
)

func installSynergyCore() {
	// Método 1: Descargar y extraer .tar.gz
	synergycore_tar_url := "https://github.com/symless/synergy-core/archive/refs/tags/1.14.6.19-stable.tar.gz"
	synergycore_cmd_tar := exec.Command("curl", "-L", synergycore_tar_url, "-o", "package.tar.gz")
	err := synergycore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	synergycore_zip_url := "https://github.com/symless/synergy-core/archive/refs/tags/1.14.6.19-stable.zip"
	synergycore_cmd_zip := exec.Command("curl", "-L", synergycore_zip_url, "-o", "package.zip")
	err = synergycore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	synergycore_bin_url := "https://github.com/symless/synergy-core/archive/refs/tags/1.14.6.19-stable.bin"
	synergycore_cmd_bin := exec.Command("curl", "-L", synergycore_bin_url, "-o", "binary.bin")
	err = synergycore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	synergycore_src_url := "https://github.com/symless/synergy-core/archive/refs/tags/1.14.6.19-stable.src.tar.gz"
	synergycore_cmd_src := exec.Command("curl", "-L", synergycore_src_url, "-o", "source.tar.gz")
	err = synergycore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	synergycore_cmd_direct := exec.Command("./binary")
	err = synergycore_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pugixml")
	exec.Command("latte", "install", "pugixml").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libnotify")
	exec.Command("latte", "install", "libnotify").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: libxinerama")
	exec.Command("latte", "install", "libxinerama").Run()
	fmt.Println("Instalando dependencia: libxkbfile")
	exec.Command("latte", "install", "libxkbfile").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: libxtst")
	exec.Command("latte", "install", "libxtst").Run()
}
