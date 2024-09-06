package main

// Mvt - Mobile device forensic toolkit
// Homepage: https://docs.mvt.re/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installMvt() {
	// Método 1: Descargar y extraer .tar.gz
	mvt_tar_url := "https://files.pythonhosted.org/packages/0d/12/a87132ab005aaa685663348df0a927c123a921ad8385813c83098c544269/mvt-2.5.4.tar.gz"
	mvt_cmd_tar := exec.Command("curl", "-L", mvt_tar_url, "-o", "package.tar.gz")
	err := mvt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mvt_zip_url := "https://files.pythonhosted.org/packages/0d/12/a87132ab005aaa685663348df0a927c123a921ad8385813c83098c544269/mvt-2.5.4.zip"
	mvt_cmd_zip := exec.Command("curl", "-L", mvt_zip_url, "-o", "package.zip")
	err = mvt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mvt_bin_url := "https://files.pythonhosted.org/packages/0d/12/a87132ab005aaa685663348df0a927c123a921ad8385813c83098c544269/mvt-2.5.4.bin"
	mvt_cmd_bin := exec.Command("curl", "-L", mvt_bin_url, "-o", "binary.bin")
	err = mvt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mvt_src_url := "https://files.pythonhosted.org/packages/0d/12/a87132ab005aaa685663348df0a927c123a921ad8385813c83098c544269/mvt-2.5.4.src.tar.gz"
	mvt_cmd_src := exec.Command("curl", "-L", mvt_src_url, "-o", "source.tar.gz")
	err = mvt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mvt_cmd_direct := exec.Command("./binary")
	err = mvt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
