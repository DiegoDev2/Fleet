package main

// SvtplayDl - Download videos from https://www.svtplay.se/
// Homepage: https://svtplay-dl.se/

import (
	"fmt"
	
	"os/exec"
)

func installSvtplayDl() {
	// Método 1: Descargar y extraer .tar.gz
	svtplaydl_tar_url := "https://files.pythonhosted.org/packages/69/7f/5a513d0a4cbf405ecb4657df6e268798d2cbcb9fac51b3a931d88061d330/svtplay_dl-4.97.1.tar.gz"
	svtplaydl_cmd_tar := exec.Command("curl", "-L", svtplaydl_tar_url, "-o", "package.tar.gz")
	err := svtplaydl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svtplaydl_zip_url := "https://files.pythonhosted.org/packages/69/7f/5a513d0a4cbf405ecb4657df6e268798d2cbcb9fac51b3a931d88061d330/svtplay_dl-4.97.1.zip"
	svtplaydl_cmd_zip := exec.Command("curl", "-L", svtplaydl_zip_url, "-o", "package.zip")
	err = svtplaydl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svtplaydl_bin_url := "https://files.pythonhosted.org/packages/69/7f/5a513d0a4cbf405ecb4657df6e268798d2cbcb9fac51b3a931d88061d330/svtplay_dl-4.97.1.bin"
	svtplaydl_cmd_bin := exec.Command("curl", "-L", svtplaydl_bin_url, "-o", "binary.bin")
	err = svtplaydl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svtplaydl_src_url := "https://files.pythonhosted.org/packages/69/7f/5a513d0a4cbf405ecb4657df6e268798d2cbcb9fac51b3a931d88061d330/svtplay_dl-4.97.1.src.tar.gz"
	svtplaydl_cmd_src := exec.Command("curl", "-L", svtplaydl_src_url, "-o", "source.tar.gz")
	err = svtplaydl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svtplaydl_cmd_direct := exec.Command("./binary")
	err = svtplaydl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
