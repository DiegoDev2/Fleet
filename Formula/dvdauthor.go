package main

// Dvdauthor - DVD-authoring toolset
// Homepage: https://dvdauthor.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDvdauthor() {
	// Método 1: Descargar y extraer .tar.gz
	dvdauthor_tar_url := "https://downloads.sourceforge.net/project/dvdauthor/dvdauthor-0.7.2.tar.gz"
	dvdauthor_cmd_tar := exec.Command("curl", "-L", dvdauthor_tar_url, "-o", "package.tar.gz")
	err := dvdauthor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvdauthor_zip_url := "https://downloads.sourceforge.net/project/dvdauthor/dvdauthor-0.7.2.zip"
	dvdauthor_cmd_zip := exec.Command("curl", "-L", dvdauthor_zip_url, "-o", "package.zip")
	err = dvdauthor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvdauthor_bin_url := "https://downloads.sourceforge.net/project/dvdauthor/dvdauthor-0.7.2.bin"
	dvdauthor_cmd_bin := exec.Command("curl", "-L", dvdauthor_bin_url, "-o", "binary.bin")
	err = dvdauthor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvdauthor_src_url := "https://downloads.sourceforge.net/project/dvdauthor/dvdauthor-0.7.2.src.tar.gz"
	dvdauthor_cmd_src := exec.Command("curl", "-L", dvdauthor_src_url, "-o", "source.tar.gz")
	err = dvdauthor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvdauthor_cmd_direct := exec.Command("./binary")
	err = dvdauthor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libdvdread")
	exec.Command("latte", "install", "libdvdread").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
