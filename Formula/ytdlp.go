package main

// YtDlp - Feature-rich command-line audio/video downloader
// Homepage: https://github.com/yt-dlp/yt-dlp

import (
	"fmt"
	
	"os/exec"
)

func installYtDlp() {
	// Método 1: Descargar y extraer .tar.gz
	ytdlp_tar_url := "https://files.pythonhosted.org/packages/d8/28/83ec43b75afd9e9840680757000fc75e68d3d221621090b3ca7601ca8129/yt_dlp-2024.8.6.tar.gz"
	ytdlp_cmd_tar := exec.Command("curl", "-L", ytdlp_tar_url, "-o", "package.tar.gz")
	err := ytdlp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ytdlp_zip_url := "https://files.pythonhosted.org/packages/d8/28/83ec43b75afd9e9840680757000fc75e68d3d221621090b3ca7601ca8129/yt_dlp-2024.8.6.zip"
	ytdlp_cmd_zip := exec.Command("curl", "-L", ytdlp_zip_url, "-o", "package.zip")
	err = ytdlp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ytdlp_bin_url := "https://files.pythonhosted.org/packages/d8/28/83ec43b75afd9e9840680757000fc75e68d3d221621090b3ca7601ca8129/yt_dlp-2024.8.6.bin"
	ytdlp_cmd_bin := exec.Command("curl", "-L", ytdlp_bin_url, "-o", "binary.bin")
	err = ytdlp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ytdlp_src_url := "https://files.pythonhosted.org/packages/d8/28/83ec43b75afd9e9840680757000fc75e68d3d221621090b3ca7601ca8129/yt_dlp-2024.8.6.src.tar.gz"
	ytdlp_cmd_src := exec.Command("curl", "-L", ytdlp_src_url, "-o", "source.tar.gz")
	err = ytdlp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ytdlp_cmd_direct := exec.Command("./binary")
	err = ytdlp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: make")
	exec.Command("latte", "install", "make").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
