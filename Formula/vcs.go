package main

// Vcs - Creates video contact sheets (previews) of videos
// Homepage: https://p.outlyer.net/vcs/

import (
	"fmt"
	
	"os/exec"
)

func installVcs() {
	// Método 1: Descargar y extraer .tar.gz
	vcs_tar_url := "https://p.outlyer.net/files/vcs/vcs-1.13.4.tar.gz"
	vcs_cmd_tar := exec.Command("curl", "-L", vcs_tar_url, "-o", "package.tar.gz")
	err := vcs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vcs_zip_url := "https://p.outlyer.net/files/vcs/vcs-1.13.4.zip"
	vcs_cmd_zip := exec.Command("curl", "-L", vcs_zip_url, "-o", "package.zip")
	err = vcs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vcs_bin_url := "https://p.outlyer.net/files/vcs/vcs-1.13.4.bin"
	vcs_cmd_bin := exec.Command("curl", "-L", vcs_bin_url, "-o", "binary.bin")
	err = vcs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vcs_src_url := "https://p.outlyer.net/files/vcs/vcs-1.13.4.src.tar.gz"
	vcs_cmd_src := exec.Command("curl", "-L", vcs_src_url, "-o", "source.tar.gz")
	err = vcs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vcs_cmd_direct := exec.Command("./binary")
	err = vcs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: gnu-getopt")
	exec.Command("latte", "install", "gnu-getopt").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
}
