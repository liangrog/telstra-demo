package telstra;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import telstra.Movie;
import telstra.MovieRepository;

@Controller  
public class MainController {
	@Autowired

	private MovieRepository movieRepository;

	@GetMapping(path="/")
	public @ResponseBody Iterable<Movie> getAllMovies() {
		// This returns a JSON or XML with the users
		return movieRepository.findAll();
	}
}
