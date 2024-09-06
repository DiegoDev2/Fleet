package main

// BandcampDl - Simple python script to download Bandcamp albums
// Homepage: https://github.com/iheanyi/bandcamp-dl

import (
	"fmt"
	
	"os/exec"
)

func installBandcampDl() {
	// Método 1: Descargar y extraer .tar.gz
	bandcampdl_tar_url := "https://files.pythonhosted.org/packages/e5/4d/d463bcc20602f5385e0441cd7171b1fe6b67e2bb76240ae0f2684de6c022/bandcamp-downloader-0.0.15.tar.gz"
	bandcampdl_cmd_tar := exec.Command("curl", "-L", bandcampdl_tar_url, "-o", "package.tar.gz")
	err := bandcampdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bandcampdl_zip_url := "https://files.pythonhosted.org/packages/e5/4d/d463bcc20602f5385e0441cd7171b1fe6b67e2bb76240ae0f2684de6c022/bandcamp-downloader-0.0.15.zip"
	bandcampdl_cmd_zip := exec.Command("curl", "-L", bandcampdl_zip_url, "-o", "package.zip")
	err = bandcampdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bandcampdl_bin_url := "https://files.pythonhosted.org/packages/e5/4d/d463bcc20602f5385e0441cd7171b1fe6b67e2bb76240ae0f2684de6c022/bandcamp-downloader-0.0.15.bin"
	bandcampdl_cmd_bin := exec.Command("curl", "-L", bandcampdl_bin_url, "-o", "binary.bin")
	err = bandcampdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bandcampdl_src_url := "https://files.pythonhosted.org/packages/e5/4d/d463bcc20602f5385e0441cd7171b1fe6b67e2bb76240ae0f2684de6c022/bandcamp-downloader-0.0.15.src.tar.gz"
	bandcampdl_cmd_src := exec.Command("curl", "-L", bandcampdl_src_url, "-o", "source.tar.gz")
	err = bandcampdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bandcampdl_cmd_direct := exec.Command("./binary")
	err = bandcampdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
