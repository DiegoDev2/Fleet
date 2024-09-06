package main

// Wdfs - Webdav file system
// Homepage: http://noedler.de/projekte/wdfs/

import (
	"fmt"
	
	"os/exec"
)

func installWdfs() {
	// Método 1: Descargar y extraer .tar.gz
	wdfs_tar_url := "http://noedler.de/projekte/wdfs/wdfs-1.4.2.tar.gz"
	wdfs_cmd_tar := exec.Command("curl", "-L", wdfs_tar_url, "-o", "package.tar.gz")
	err := wdfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wdfs_zip_url := "http://noedler.de/projekte/wdfs/wdfs-1.4.2.zip"
	wdfs_cmd_zip := exec.Command("curl", "-L", wdfs_zip_url, "-o", "package.zip")
	err = wdfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wdfs_bin_url := "http://noedler.de/projekte/wdfs/wdfs-1.4.2.bin"
	wdfs_cmd_bin := exec.Command("curl", "-L", wdfs_bin_url, "-o", "binary.bin")
	err = wdfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wdfs_src_url := "http://noedler.de/projekte/wdfs/wdfs-1.4.2.src.tar.gz"
	wdfs_cmd_src := exec.Command("curl", "-L", wdfs_src_url, "-o", "source.tar.gz")
	err = wdfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wdfs_cmd_direct := exec.Command("./binary")
	err = wdfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: neon")
	exec.Command("latte", "install", "neon").Run()
}
