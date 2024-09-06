package main

// ArxivLatexCleaner - Clean LaTeX code to submit to arXiv
// Homepage: https://github.com/google-research/arxiv-latex-cleaner

import (
	"fmt"
	
	"os/exec"
)

func installArxivLatexCleaner() {
	// Método 1: Descargar y extraer .tar.gz
	arxivlatexcleaner_tar_url := "https://files.pythonhosted.org/packages/7b/be/e0afb37ba09060368e3858c8248328faf187d814f9cb9da00e5611d150d0/arxiv_latex_cleaner-1.0.8.tar.gz"
	arxivlatexcleaner_cmd_tar := exec.Command("curl", "-L", arxivlatexcleaner_tar_url, "-o", "package.tar.gz")
	err := arxivlatexcleaner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arxivlatexcleaner_zip_url := "https://files.pythonhosted.org/packages/7b/be/e0afb37ba09060368e3858c8248328faf187d814f9cb9da00e5611d150d0/arxiv_latex_cleaner-1.0.8.zip"
	arxivlatexcleaner_cmd_zip := exec.Command("curl", "-L", arxivlatexcleaner_zip_url, "-o", "package.zip")
	err = arxivlatexcleaner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arxivlatexcleaner_bin_url := "https://files.pythonhosted.org/packages/7b/be/e0afb37ba09060368e3858c8248328faf187d814f9cb9da00e5611d150d0/arxiv_latex_cleaner-1.0.8.bin"
	arxivlatexcleaner_cmd_bin := exec.Command("curl", "-L", arxivlatexcleaner_bin_url, "-o", "binary.bin")
	err = arxivlatexcleaner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arxivlatexcleaner_src_url := "https://files.pythonhosted.org/packages/7b/be/e0afb37ba09060368e3858c8248328faf187d814f9cb9da00e5611d150d0/arxiv_latex_cleaner-1.0.8.src.tar.gz"
	arxivlatexcleaner_cmd_src := exec.Command("curl", "-L", arxivlatexcleaner_src_url, "-o", "source.tar.gz")
	err = arxivlatexcleaner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arxivlatexcleaner_cmd_direct := exec.Command("./binary")
	err = arxivlatexcleaner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
