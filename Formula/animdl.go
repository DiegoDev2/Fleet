package main

// Animdl - Anime downloader and streamer
// Homepage: https://github.com/justfoolingaround/animdl

import (
	"fmt"
	
	"os/exec"
)

func installAnimdl() {
	// Método 1: Descargar y extraer .tar.gz
	animdl_tar_url := "https://files.pythonhosted.org/packages/5b/79/4be6ac2caca32dea6fe500e5f5df9d74a3a5ce1d500175c3a7b69500bb3f/animdl-1.7.27.tar.gz"
	animdl_cmd_tar := exec.Command("curl", "-L", animdl_tar_url, "-o", "package.tar.gz")
	err := animdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	animdl_zip_url := "https://files.pythonhosted.org/packages/5b/79/4be6ac2caca32dea6fe500e5f5df9d74a3a5ce1d500175c3a7b69500bb3f/animdl-1.7.27.zip"
	animdl_cmd_zip := exec.Command("curl", "-L", animdl_zip_url, "-o", "package.zip")
	err = animdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	animdl_bin_url := "https://files.pythonhosted.org/packages/5b/79/4be6ac2caca32dea6fe500e5f5df9d74a3a5ce1d500175c3a7b69500bb3f/animdl-1.7.27.bin"
	animdl_cmd_bin := exec.Command("curl", "-L", animdl_bin_url, "-o", "binary.bin")
	err = animdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	animdl_src_url := "https://files.pythonhosted.org/packages/5b/79/4be6ac2caca32dea6fe500e5f5df9d74a3a5ce1d500175c3a7b69500bb3f/animdl-1.7.27.src.tar.gz"
	animdl_cmd_src := exec.Command("curl", "-L", animdl_src_url, "-o", "source.tar.gz")
	err = animdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	animdl_cmd_direct := exec.Command("./binary")
	err = animdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
