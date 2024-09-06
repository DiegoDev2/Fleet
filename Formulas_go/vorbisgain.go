package main

// Vorbisgain - Add Replay Gain volume tags to Ogg Vorbis files
// Homepage: https://sjeng.org/vorbisgain.html

import (
	"fmt"
	
	"os/exec"
)

func installVorbisgain() {
	// Método 1: Descargar y extraer .tar.gz
	vorbisgain_tar_url := "https://sjeng.org/ftp/vorbis/vorbisgain-0.37.tar.gz"
	vorbisgain_cmd_tar := exec.Command("curl", "-L", vorbisgain_tar_url, "-o", "package.tar.gz")
	err := vorbisgain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vorbisgain_zip_url := "https://sjeng.org/ftp/vorbis/vorbisgain-0.37.zip"
	vorbisgain_cmd_zip := exec.Command("curl", "-L", vorbisgain_zip_url, "-o", "package.zip")
	err = vorbisgain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vorbisgain_bin_url := "https://sjeng.org/ftp/vorbis/vorbisgain-0.37.bin"
	vorbisgain_cmd_bin := exec.Command("curl", "-L", vorbisgain_bin_url, "-o", "binary.bin")
	err = vorbisgain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vorbisgain_src_url := "https://sjeng.org/ftp/vorbis/vorbisgain-0.37.src.tar.gz"
	vorbisgain_cmd_src := exec.Command("curl", "-L", vorbisgain_src_url, "-o", "source.tar.gz")
	err = vorbisgain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vorbisgain_cmd_direct := exec.Command("./binary")
	err = vorbisgain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
}
