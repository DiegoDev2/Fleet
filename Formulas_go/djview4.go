package main

// Djview4 - Viewer for the DjVu image format
// Homepage: https://djvu.sourceforge.net/djview4.html

import (
	"fmt"
	
	"os/exec"
)

func installDjview4() {
	// Método 1: Descargar y extraer .tar.gz
	djview4_tar_url := "https://downloads.sourceforge.net/project/djvu/DjView/4.12/djview-4.12.tar.gz"
	djview4_cmd_tar := exec.Command("curl", "-L", djview4_tar_url, "-o", "package.tar.gz")
	err := djview4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djview4_zip_url := "https://downloads.sourceforge.net/project/djvu/DjView/4.12/djview-4.12.zip"
	djview4_cmd_zip := exec.Command("curl", "-L", djview4_zip_url, "-o", "package.zip")
	err = djview4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djview4_bin_url := "https://downloads.sourceforge.net/project/djvu/DjView/4.12/djview-4.12.bin"
	djview4_cmd_bin := exec.Command("curl", "-L", djview4_bin_url, "-o", "binary.bin")
	err = djview4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djview4_src_url := "https://downloads.sourceforge.net/project/djvu/DjView/4.12/djview-4.12.src.tar.gz"
	djview4_cmd_src := exec.Command("curl", "-L", djview4_src_url, "-o", "source.tar.gz")
	err = djview4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djview4_cmd_direct := exec.Command("./binary")
	err = djview4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: djvulibre")
exec.Command("latte", "install", "djvulibre")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
