package main

// Manim - Animation engine for explanatory math videos
// Homepage: https://www.manim.community

import (
	"fmt"
	
	"os/exec"
)

func installManim() {
	// Método 1: Descargar y extraer .tar.gz
	manim_tar_url := "https://files.pythonhosted.org/packages/83/5f/717ba528eb191124211036ec710bafd605dc7f7bb948a41219a8dd1124b6/manim-0.18.1.tar.gz"
	manim_cmd_tar := exec.Command("curl", "-L", manim_tar_url, "-o", "package.tar.gz")
	err := manim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	manim_zip_url := "https://files.pythonhosted.org/packages/83/5f/717ba528eb191124211036ec710bafd605dc7f7bb948a41219a8dd1124b6/manim-0.18.1.zip"
	manim_cmd_zip := exec.Command("curl", "-L", manim_zip_url, "-o", "package.zip")
	err = manim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	manim_bin_url := "https://files.pythonhosted.org/packages/83/5f/717ba528eb191124211036ec710bafd605dc7f7bb948a41219a8dd1124b6/manim-0.18.1.bin"
	manim_cmd_bin := exec.Command("curl", "-L", manim_bin_url, "-o", "binary.bin")
	err = manim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	manim_src_url := "https://files.pythonhosted.org/packages/83/5f/717ba528eb191124211036ec710bafd605dc7f7bb948a41219a8dd1124b6/manim-0.18.1.src.tar.gz"
	manim_cmd_src := exec.Command("curl", "-L", manim_src_url, "-o", "source.tar.gz")
	err = manim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	manim_cmd_direct := exec.Command("./binary")
	err = manim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cython")
exec.Command("latte", "install", "cython")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: py3cairo")
exec.Command("latte", "install", "py3cairo")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: scipy")
exec.Command("latte", "install", "scipy")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: patchelf")
exec.Command("latte", "install", "patchelf")
}
