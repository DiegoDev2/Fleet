package formulas

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallNmap() {
	switch runtime.GOOS {
	case "darwin":
		installNmapMac()
	default:
		fmt.Println("Este script solo está preparado para macOS (darwin).")
	}
}

func installNmapMac() {
	url := "https://nmap.org/dist/nmap-7.95.tar.bz2"

	fmt.Println("Descargando Nmap...")
	download := exec.Command("curl", "-L", url, "-o", "nmap.tar.bz2")
	download.Stdout = os.Stdout
	download.Stderr = os.Stderr
	if err := download.Run(); err != nil {
		fmt.Println("Error descargando Nmap:", err)
		return
	}

	fmt.Println("Descomprimiendo Nmap...")
	extract := exec.Command("tar", "-xjf", "nmap.tar.bz2")
	extract.Stdout = os.Stdout
	extract.Stderr = os.Stderr
	if err := extract.Run(); err != nil {
		fmt.Println("Error descomprimiendo Nmap:", err)
		return
	}

	fmt.Println("Entrando en el directorio de Nmap...")
	if err := os.Chdir("nmap-7.95"); err != nil {
		fmt.Println("Error cambiando de directorio:", err)
		return
	}

	fmt.Println("Compilando Nmap...")
	configure := exec.Command("./configure")
	configure.Stdout = os.Stdout
	configure.Stderr = os.Stderr
	if err := configure.Run(); err != nil {
		fmt.Println("Error configurando Nmap:", err)
		return
	}

	make := exec.Command("make")
	make.Stdout = os.Stdout
	make.Stderr = os.Stderr
	if err := make.Run(); err != nil {
		fmt.Println("Error compilando Nmap:", err)
		return
	}

	fmt.Println("Instalando Nmap...")
	install := exec.Command("sudo", "make", "install")
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	if err := install.Run(); err != nil {
		fmt.Println("Error instalando Nmap:", err)
		return
	}

	fmt.Println("Nmap instalado correctamente.")
}
